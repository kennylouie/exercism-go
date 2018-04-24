package tree

import (
  "errors"
  "sort"
  "sync"
)

type Record struct {
  ID, Parent int
}

type Node struct {
  ID  int
  Children  []*Node
}

// main tree building function
func Build(records []Record) (*Node, error) {

  // initial empty records check
  if len(records) == 0 {
    return nil, nil
  }

  // error channel
  errorChan := make(chan error, 10)

  // concurrently check for different errors while also building tree
  go errorsRoot(records, errorChan)
  go errorsContinuity(records, errorChan)
  go errorsDirectCycling(records, errorChan)
  go errorsIndirectCycling(records, errorChan)

  // channels for building trees
  todoChan := make(chan *Node, 10)
  toOrderChan := make(chan *Node, 10)

  // init with a root
  root := &Node{ID: 0}
  todoChan<- root
  // build tree
  var wg sync.WaitGroup
  wg.Add(2 * len(records))

  // concurrently find children and organize children
  go AllocateChildrenFinders(records, todoChan, toOrderChan, errorChan, &wg)
  go AllocateChildrenOrganizers(records, toOrderChan, todoChan, errorChan, &wg)

  for i := 0; i < 2 * len(records) + 2; i++ {
    if e := <-errorChan; e != nil {
      return nil, e
    }
  }

  wg.Wait()
  return root, nil
}


// creates concurrent workers to find children
func AllocateChildrenFinders(records []Record, todoChan <-chan *Node, toOrderChan chan<- *Node, errorChan chan<- error, wg *sync.WaitGroup) {
  for range records {
    go FindChildren(records, todoChan, toOrderChan, errorChan, wg)
  }
}

// children finder worker
func FindChildren(records []Record, todoChan <-chan *Node, toOrderChan chan<- *Node, errorChan chan<- error, wg *sync.WaitGroup) {
  n := <-todoChan

  for _, r := range records {
    if r.Parent == n.ID {

    // check for errors
    if r.ID < n.ID {
      errorChan<- errors.New("Cannot have an ID lower than that of parent")
    } else if r.ID == n.ID {
      if r.ID != 0 {
        errorChan<- errors.New("Only instance when record ID is the same as parent ID is when ID == 0")
      }
    } else {

      // build children
      newNode := &Node{ID: r.ID}
      n.Children = append(n.Children, newNode)

      }
    }
  }

  toOrderChan<- n
  errorChan<- nil
  wg.Done()
}

// creates concurrent workers to organize children slices
func AllocateChildrenOrganizers(records []Record, toOrderChan chan *Node, todoChan chan *Node, errorChan chan<- error, wg *sync.WaitGroup) {
  for range records {
    go OrderChildren(records, toOrderChan, todoChan, errorChan, wg)
  }
}

// children organizer worker
func OrderChildren(records []Record, toOrderChan chan *Node, todoChan chan *Node, errorChan chan<- error, wg *sync.WaitGroup) {

  n := <-toOrderChan

  if len(n.Children) > 1 {
    sort.Slice(n.Children, func(i, j int) bool { return n.Children[i].ID < n.Children[j].ID })

    // check for repeat elements
    sliceCheck := make(map[int]bool)

    for _, c := range n.Children {
      if sliceCheck[c.ID] {
        errorChan<- errors.New("duplicate node")
      }
      sliceCheck[c.ID] = true
    }

  // put records to be searched into todochannel
    for i := range n.Children {
      todoChan<- n.Children[i]
    }

  }

  if len(n.Children) == 1 {
    todoChan<- n.Children[0]
  }

  errorChan<- nil
  wg.Done()
}

// error checking goroutines
// check for a single root
func errorsRoot(records []Record, errorChan chan<- error) {

  n := 0
  for _, r := range records {
    if r.ID == 0 {
      n += 1
      if r.Parent != 0 {
        errorChan<- errors.New("root node has parent")
      }
    }
  }

  if n != 1 {
      errorChan<- errors.New("duplicate root")
  }

  errorChan<- nil
}

// check for non-continuality
func errorsContinuity(records []Record, errorChan chan<- error) {

  for _, v := range records {
    if v.ID >= len(records) {
      errorChan<- errors.New("non-continous")
    }
  }

  errorChan<- nil
}

// check for direct cycling
func errorsDirectCycling(records []Record, errorChan chan<- error) {

  for _, v := range records {
    if v.ID == v.Parent {
      if v.ID != 0 {
        errorChan<- errors.New("cycle directly")
      }
    }
  }

  errorChan<- nil
}

// check for indirect cycling
func errorsIndirectCycling(records []Record, errorChan chan<- error) {
  max := 0

  for _, v := range records {
    if v.Parent > max {
      max = v.Parent
    }
  }

  for i := 0; i <= max; i++ {

    exist := false
    for _, v := range records {
      if v.Parent == i {
        exist = true
      }
    }

    if !exist {
      errorChan<- errors.New("cycle indirectly")
    }

  }

  errorChan<- nil
}

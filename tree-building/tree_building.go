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
	
	// empty records check
	if len(records) == 0 {
		return nil, nil
	}
	  
	// channels
	todoChan := make(chan *Node, 10)
  toOrderChan := make(chan *Node, 10)
  errorChan := make(chan error, 10)
	
	// check for a single root
	n := 0
	for _, r := range records {
	  if r.ID == 0 {
	    n += 1
	    if r.Parent != 0 {
	      return nil, errors.New("root node has parent")
	    }
	  }
	}
	if n != 1 {
	    return nil, errors.New("duplicate root")
	}
	
	// check for non-continuality
	for _, v := range records {
	  if v.ID >= len(records) {
	    return nil, errors.New("non-continous")
	  }
	}
	
	// check for direct cycling
	for _, v := range records {
	  if v.ID == v.Parent {
	    if v.ID != 0 {
	      return nil, errors.New("cycle directly")
	    }
	  }
	}
	
	// check for indirect cycling
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
	    return nil, errors.New("cycle indirectly")
	  }
	}
	
	// init with a root
	root := &Node{ID: 0}
	todoChan<- root
	
	// build tree
	var wg sync.WaitGroup
	wg.Add(2 * len(records))
	
	for range records {
	 
	  r := <-todoChan
	  go r.AddChildren(records, &wg, errorChan, toOrderChan)
	  
	  c := <-toOrderChan
	  go c.OrderChildren(&wg, todoChan, errorChan)
	  
	  if e := <-errorChan; e != nil {
	    return nil, e
	  }
	  
	}
	
	wg.Wait()
	
	return root, nil
}


// checks records for children
func (n *Node) AddChildren(records []Record, wg *sync.WaitGroup, errorChan chan<- error, toOrderChan chan *Node) {
  
  for i := 0; i < len(records); i++ {
    
    if records[i].Parent == n.ID {
      
      // check for errors
      if records[i].ID < n.ID {
        errorChan<- errors.New("Cannot have an ID lower than that of parent")
      } else if records[i].ID == n.ID {
        if records[i].ID != 0 {
          errorChan<- errors.New("Only instance when record ID is the same as parent ID is when ID == 0")
        }
      } else {
        
        // build children
        newNode := &Node{ID: records[i].ID}
        n.Children = append(n.Children, newNode)
        
      }
    }
  }
  
  toOrderChan<- n
  errorChan<- nil
  
  wg.Done()
}

// organize children
func (n *Node) OrderChildren(wg *sync.WaitGroup, todoChan chan *Node, errorChan chan<- error) {
  
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
  
  wg.Done()
}

package letter

import "sync"

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
  m := FreqMap{}
  for _, r := range s {
    m[r]++
  }
  return m
}

type muxMap struct {
  m FreqMap
  mux sync.Mutex
}

func ConcurrentFrequency(s []string) FreqMap {

  mapChan := make(chan FreqMap, len(s))
  var wg sync.WaitGroup
  m := muxMap{m: make(FreqMap)}

  for _, v := range s {
    wg.Add(1)
    go UnitFrequency(v, mapChan, &wg)
  }

  for range s {
    wg.Add(1)
    go m.AddToMap(mapChan, &wg)
  }

  wg.Wait()

  return m.m

}

func UnitFrequency(s string, mapChan chan<- FreqMap, wg *sync.WaitGroup) {
  mapChan<- Frequency(s)
  wg.Done()
}

func (m *muxMap) AddToMap(mapChan <-chan FreqMap, wg *sync.WaitGroup) {
  aMap := <-mapChan
  for k, v := range aMap {
    m.mux.Lock()
    m.m[k] += v
    m.mux.Unlock()
  }
  wg.Done()
}


// Vince's code
// package letter

// //import "fmt"
// import "sync"

// type Occurrence map[rune]int

// func Frequency(input string) Occurrence {
//     dict := Occurrence{}
//     for _, r := range input {
//         dict[r]++
//     }
//     return dict
// }

// func Concurrent(message string, c chan Occurrence, wg *sync.WaitGroup) {
//     dict := Occurrence{}
//     for _, r := range message {
//         dict[r]++
//     }
//     c <- dict
//     wg.Done()
// }

// func ConcurrentFrequency(input []string) Occurrence {
//     var wg sync.WaitGroup
//     dict := Occurrence{}
//     freqChan := make(chan Occurrence, len(input))

//     wg.Add(len(input))
//     for _, n := range input {
//         go Concurrent(n, freqChan, &wg)
//     }
//     wg.Wait()

//     for i := 0; i < len(input); i++ {
//         for k, v := range <-freqChan {
//             dict[k] += v
//         }
//     }

//     return dict
// }

package clock

import (
  "fmt"
)

const (
  minutesInHour = 60
  hoursInDay = 24
)

type Clock struct {
  h, m int
}

func (c *Clock) String() string {
  return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func New(h, m int) Clock {

  minutes := h*minutesInHour + m

  clock := Clock{0, 0}.Add(minutes)

  return clock
}

func (c Clock) Add(min int) Clock {

  // calculate time in minutes
  mins := min + c.h*minutesInHour + c.m

  newHour := mins / minutesInHour

  // edge case when mins are a negative fraction
  if mins % minutesInHour < 0 {
    newHour -= 1
  }

  for newHour >= hoursInDay {
    newHour -= hoursInDay
  }

  for newHour < 0 {
    newHour += hoursInDay
  }

  for mins < 0 {
    mins += minutesInHour
  }

  c.h = newHour
  c.m = mins % minutesInHour

  return c

}

func (c Clock) Subtract(min int) Clock {
  c = c.Add(-min)
  return c
}

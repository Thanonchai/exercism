package clock

import "fmt"

type Clock struct {
	Minute int
}

const TotalMinutes int = 1440

func New(hour, minute int) Clock {
	clock := Clock{0}
	clock.Minute += hour * 60
	clock.Minute += minute
	clock.afterModify()
	return clock
}

func (c Clock) Add(minutes int) Clock {
	c.Minute += minutes
	c.afterModify()
	return c
}

func (c Clock) Subtract(minutes int) Clock {
	c.Minute -= minutes
	c.afterModify()
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Minute/60, c.Minute%60)
}

func (c *Clock) afterModify() {
	c.Minute %= TotalMinutes
	c.Minute += TotalMinutes
	c.Minute %= TotalMinutes
}

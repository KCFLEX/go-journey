// Package dog provides a function that turns human years to dogs years
package dog

// func Years takes in a parameter type int which
// represents the humans age and returns an int
// which represents the dogs age
func Years(hyears int) int {
	dogyears := 7 * hyears
	return dogyears
}

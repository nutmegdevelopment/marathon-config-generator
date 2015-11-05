package main

import "fmt"

// Define a stringslice type to hold the key/value pairs passed in via the var
// command line flag.
type stringslice []string

// Now implement the two methods for the flag.Value interface:

// The first method is String() string
func (s *stringslice) String() string {
	return fmt.Sprintf("%s", *s)
}

// The second method is Set(value string) error
func (s *stringslice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

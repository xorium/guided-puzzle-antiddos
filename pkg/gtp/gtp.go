// Package gtp (Guided Tour Puzzles) provides primitive data structures and
// algorithms to operating with them.
//
// When a server_guide suspects that it is under attack or its load is
// above a certain threshold, it asks all clients to solve a puzzle
// prior to receiving service. In the GTP protocol, the puzzle
// is simply a tour that needs to be completed by the client
// via taking round-trips to a set of special nodes, called *tour
// guides*, in a sequential order.
//
// We call this tour a *guided tour*, because the client should not
// know the order of the tour a priori, and each tour guide must direct
// the client towards the next tour guide. Each tour guide may appear zero or
// more times in a tour.
//
// And the term *stop* is used to represent a single appearance of a
// tour guide in a given tour.

package gtp

/*
Package pocketlog exposes an API to log your work.

First instantiate a logger with pocketlog.New, and giving it a threshold
Messages of lesser criticality won't be logged.

Sharing the logger is the responsibility of the caller.

The logger can be called to tog messages on three levels:
  - Debug: mostly used to debug codes, follow step-by-step process
  - Info: valuable messages providing insights to the milestones
  - Error: error messages to understand what went wrong
*/
package pocketlog

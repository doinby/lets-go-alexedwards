// Model/DB-specific Error Handler

package models

import "errors"

var ErrNoRecord = errors.New("models: no matching error record found")

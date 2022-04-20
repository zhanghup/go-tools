package xls

const (
	// ContinueColumnMerged marks a continuation column within a merged cell.
	ContinueColumnMerged = "→"
	// EndColumnMerged marks the last column of a merged cell.
	EndColumnMerged = "⇥"

	// ContinueRowMerged marks a continuation row within a merged cell.
	ContinueRowMerged = "↓"
	// EndRowMerged marks the last row of a merged cell.
	EndRowMerged = "⤓"
)

type Source interface {
	// List the individual data tables within this source.
	List() ([]string, error)

	// Get a Collection from the source by name.
	Get(name string) (Collection, error)

	// Close the source and discard memory.
	Close() error
}

// Collection represents an iterable collection of records.
type Collection interface {
	// Next advances to the next record of content.
	// It MUST be called prior to any Scan().
	Next() bool

	// Strings extracts values from the current record into a list of strings.
	Strings() []string

	// Types extracts the data types from the current record into a list.
	// options: "boolean", "integer", "float", "string", "date",
	// and special cases: "blank", "hyperlink" which are string types
	Types() []string

	// Formats extracts the format codes for the current record into a list.
	Formats() []string

	// Scan extracts values from the current record into the provided arguments
	// Arguments must be pointers to one of 5 supported types:
	//     bool, int64, float64, string, or time.Time
	// If invalid, returns ErrInvalidScanType
	Scan(args ...any) error

	// IsEmpty returns true if there are no data values.
	IsEmpty() bool

	// Err returns the last error that occured.
	Err() error
}

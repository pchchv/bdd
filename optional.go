package bdd

const (
	// Formatted option
	_                  option = 1 << (31 - iota)
	CanDefaultString          // can use .(fmt.Stringer)
	CanFilterDuplicate        // Filter duplicates
	CanRowSpan                // Fold line

	// Formatted style
	_           style = iota
	StylePrint        // Display data; string without quotes
	StyleBPrint       // Display data
	StyleTPrint       // Display type and data
	StyleJPrint       // The json style display; Do not show private
)

type style int

type option uint32

func (t option) IsCanDefaultString() bool {
	return (t & CanDefaultString) != 0
}

func (t option) IsCanFilterDuplicate() bool {
	return (t & CanFilterDuplicate) != 0
}

func (t option) IsCanRowSpan() bool {
	return (t & CanRowSpan) != 0
}

package board

import (
	"reflect"
	"testing"
)

func TestColumnForPosition(t *testing.T) {
	// tab := [TOT_SIZE]int{}

	// 0-indexed columns
	tables := []struct {
		position       int
		expectedColumn int
	}{
		{0, 0},
		{1, 1},
		{19, 0},
		{20, 1},
		{21, 2},
		{18, 18},
		{38, 0},
		{37, 18},
	}

	for _, table := range tables {
		actual := GetColumnForPosition(table.position)
		if actual != table.expectedColumn {
			t.Errorf("⛔️ Position: %v, expect column: %v, got %v", table.position, table.expectedColumn, actual)
		}
	}
}

func TestRowForPosition(t *testing.T) {
	// tab := [TOT_SIZE]int{}

	// 0-indexed rows
	tables := []struct {
		position    int
		expectedRow int
	}{
		{0, 0},
		{1, 0},
		{19, 1},
		{20, 1},
		{21, 1},
		{18, 0},
		{38, 2},
		{37, 1},
	}

	for _, table := range tables {
		actual := GetRowForPosition(table.position)
		if actual != table.expectedRow {
			t.Errorf("⛔️ Position: %v, expect Row: %v, got %v", table.position, table.expectedRow, actual)
		}
	}
}

func TestGetIndexNWSEForPosition(t *testing.T) {
	// tab := [TOT_SIZE]int{}

	// 1-indexed diagonals
	tables := []struct {
		position  int
		expectedD int
	}{
		{342, 1},
		{323, 2},
		{343, 2},
		{266, 5},
		{286, 5},
		{0, 19},
		{40, 19},
		{1, 20},
		{341, 20},
	}

	for _, table := range tables {
		actual := GetIndexNWSEForPosition(table.position)
		if actual != table.expectedD {
			t.Errorf("⛔️ Position: %v, expect d: %v, got %v", table.position, table.expectedD, actual)
		}
	}
}

func TestGetIndexNESWForPosition(t *testing.T) {
	// tab := [TOT_SIZE]int{}

	// 1-indexed diagonals
	tables := []struct {
		position  int
		expectedD int
	}{
		{0, 1},
		{1, 2},
		{19, 2},
		{20, 3},
		{21, 4},
		{18, 19},
		{17, 18},
		{37, 20},
		{343, 20},
		{55, 20},
	}

	for _, table := range tables {
		actual := GetIndexNESWForPosition(table.position)
		if actual != table.expectedD {
			t.Errorf("⛔️ Position: %v, expect d: %v, got %v", table.position, table.expectedD, actual)
		}
	}
}

func TestGetRowSeqForRow(t *testing.T) {
	tab := [TOT_SIZE]int{}
	for i := 0; i < TOT_SIZE; i++ {
		tab[i] = i
	}

	tables := []struct {
		row            int
		expectedRowSeq []int
	}{
		{0, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}},
		{1, []int{19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37}},
	}

	for _, table := range tables {
		actual := GetRowSeqForRow(table.row, &tab)
		if !reflect.DeepEqual(*actual, table.expectedRowSeq) {
			t.Errorf("⛔️ Row: %v, expect column: %v, got %v", table.row, table.expectedRowSeq, *actual)
		}
	}
}

func TestGetColSeqForCol(t *testing.T) {
	tab := [TOT_SIZE]int{}
	for i := 0; i < TOT_SIZE; i++ {
		tab[i] = i
	}

	tables := []struct {
		col            int
		expectedColSeq []int
	}{
		{5, []int{5, 24, 43, 62, 81, 100, 119, 138, 157, 176, 195, 214, 233, 252, 271, 290, 309, 328, 347}},
		{11, []int{11, 30, 49, 68, 87, 106, 125, 144, 163, 182, 201, 220, 239, 258, 277, 296, 315, 334, 353}},
	}

	for _, table := range tables {
		actual := GetColSeqForCol(table.col, &tab)
		if !reflect.DeepEqual(*actual, table.expectedColSeq) {
			t.Errorf("⛔️ Col: %v, expect column: %v, got %v", table.col, table.expectedColSeq, *actual)
		}
	}
}

func TestGetDiagonalNESWSequence(t *testing.T) {
	tab := [TOT_SIZE]int{}
	for i := 0; i < TOT_SIZE; i++ {
		tab[i] = i
	}

	tables := []struct {
		d               int
		expectedNESWSeq []int
	}{
		{1, []int{0}},
		{2, []int{1, 19}},
		{3, []int{2, 20, 38}},
		{4, []int{3, 21, 39, 57}},
		{37, []int{360}},
	}

	for _, table := range tables {
		actual := GetDiagonalNESWSequence(table.d, &tab)
		if !reflect.DeepEqual(*actual, table.expectedNESWSeq) {
			t.Errorf("⛔️ d: %v, expect seq: %v, got %v", table.d, table.expectedNESWSeq, *actual)
		}
	}
}

func TestGetDiagonalNWSESequence(t *testing.T) {
	tab := [TOT_SIZE]int{}
	for i := 0; i < TOT_SIZE; i++ {
		tab[i] = i
	}

	tables := []struct {
		d               int
		expectedNWSESeq []int
	}{
		{1, []int{342}},
		{2, []int{323, 343}},
		{3, []int{304, 324, 344}},
		{6, []int{247, 267, 287, 307, 327, 347}},
		{37, []int{18}},
	}

	for _, table := range tables {
		actual := GetDiagonalNWSESequence(table.d, &tab)
		if !reflect.DeepEqual(*actual, table.expectedNWSESeq) {
			t.Errorf("⛔️ d: %v, expect seq: %v, got %v", table.d, table.expectedNWSESeq, *actual)
		}
	}
}

func TestGetDiagonalNWSESequenceIdxPos(t *testing.T) {
	tab := [TOT_SIZE]int{}
	for i := 0; i < TOT_SIZE; i++ {
		tab[i] = 0
		// tab[i] = i
	}

	tables := []struct {
		d               int
		position        int
		expectedNWSESeq []int
		expectedOffset  int
	}{
		{1, 342, []int{342}, 0},
		{2, 323, []int{323, 343}, 0},
		{2, 343, []int{323, 343}, 1},
		{3, 304, []int{304, 324, 344}, 0},
		{3, 324, []int{304, 324, 344}, 1},
		{6, 287, []int{247, 267, 287, 307, 327, 347}, 2},
		{6, 347, []int{247, 267, 287, 307, 327, 347}, 5},
		// {37, []int{18}},
	}

	for _, table := range tables {
		actualSequence, actualOffset := GetDiagonalNWSESequenceIdxPos(table.d, &tab, table.position)
		if !reflect.DeepEqual(*actualSequence, table.expectedNWSESeq) {
			t.Errorf("⛔️ d: %v, expect seq: %v, got %v", table.d, table.expectedNWSESeq, *actualSequence)
		}
		if table.expectedOffset != actualOffset {
			t.Errorf("⛔️ sequence: %v, expect offset: %v, got offset %v", table.expectedNWSESeq, table.expectedOffset, actualOffset)
		}
	}
}

func TestGetOffsetAndIndexNWSEForPosition(t *testing.T) {
	tables := []struct {
		position       int
		expectedD      int
		expectedOffset int
	}{
		{342, 1, 0},
		{323, 2, 0},
		{343, 2, 1},
		{304, 3, 0},
		{324, 3, 1},
		{287, 6, 2},
		{347, 6, 5},
		{0, 19, 0},
		{1, 20, 0},
		{2, 21, 0},
	}

	for i, table := range tables {
		actualOffset, actualD := GetOffsetAndIndexNWSEForPosition(table.position)
		if table.expectedD != actualD {
			t.Errorf("⛔️ case [%d] position: %v, expect d: %v, got d %v", i, table.position, table.expectedD, actualD)
		}
		if table.expectedOffset != actualOffset {
			t.Errorf("⛔️ case [%d] position: %v, expect offset: %v, got offset %v", i, table.position, table.expectedOffset, actualOffset)
		}
	}
}

func TestGetOffsetAndIndexNESWForPosition(t *testing.T) {
	tables := []struct {
		position       int
		expectedD      int
		expectedOffset int
	}{
		{0, 1, 0},
		{1, 2, 0},
		{3, 4, 0},
		{19, 2, 1},
		{38, 3, 2},
		{20, 3, 1},
		{18, 19, 0},
		{37, 20, 0},
		{22, 5, 1},
	}

	for i, table := range tables {
		actualOffset, actualD := GetOffsetAndIndexNESWForPosition(table.position)
		if table.expectedD != actualD {
			t.Errorf("⛔️ case [%d] position: %v, expect d: %v, got d %v", i, table.position, table.expectedD, actualD)
		}
		if table.expectedOffset != actualOffset {
			t.Errorf("⛔️ case [%d] position: %v, expect offset: %v, got offset %v", i, table.position, table.expectedOffset, actualOffset)
		}
	}
}

// [   0    1    2    3    4    5,    6    7    8    9   10   11,   12   13   14   15   16   17   18]
// [  19   20   21   22   23   24,   25   26   27   28   29   30,   31   32   33   34   35   36   37]
// [  38   39   40   41   42   43,   44   45   46   47   48   49,   50   51   52   53   54   55   56]
// [  57   58   59   60   61   62,   63   64   65   66   67   68,   69   70   71   72   73   74   75]
// [  76   77   78   79   80   81,   82   83   84   85   86   87,   88   89   90   91   92   93   94]
// [  95   96   97   98   99  100,  101  102  103  104  105  106,  107  108  109  110  111  112  113]
// [ 114  115  116  117  118  119,  120  121  122  123  124  125,  126  127  128  129  130  131  132]
// [ 133  134  135  136  137  138,  139  140  141  142  143  144,  145  146  147  148  149  150  151]
// [ 152  153  154  155  156  157,  158  159  160  161  162  163,  164  165  166  167  168  169  170]
// [ 171  172  173  174  175  176,  177  178  179  180  181  182,  183  184  185  186  187  188  189]
// [ 190  191  192  193  194  195,  196  197  198  199  200  201,  202  203  204  205  206  207  208]
// [ 209  210  211  212  213  214,  215  216  217  218  219  220,  221  222  223  224  225  226  227]
// [ 228  229  230  231  232  233,  234  235  236  237  238  239,  240  241  242  243  244  245  246]
// [ 247  248  249  250  251  252,  253  254  255  256  257  258,  259  260  261  262  263  264  265]
// [ 266  267  268  269  270  271,  272  273  274  275  276  277,  278  279  280  281  282  283  284]
// [ 285  286  287  288  289  290,  291  292  293  294  295  296,  297  298  299  300  301  302  303]
// [ 304  305  306  307  308  309,  310  311  312  313  314  315,  316  317  318  319  320  321  322]
// [ 323  324  325  326  327  328,  329  330  331  332  333  334,  335  336  337  338  339  340  341]
// [ 342  343  344  345  346  347,  348  349  350  351  352  353,  354  355  356  357  358  359  360]//

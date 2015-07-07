package base

import (
	"fmt"
)

type Version struct {
	Major   int
	Minor   int
	Build   int
	Revison string
}

func (v *Version) String() string {
	if v.Revison == "" {
		return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Build)
	}

	return fmt.Sprintf("%d.%d.%d#%s", v.Major, v.Minor, v.Build, v.Revison)
}

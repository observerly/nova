/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/nova/loader
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package main

/*****************************************************************************************************************/

import (
	"fmt"
	"io"
	"os"

	gormschema "ariga.io/atlas-provider-gorm/gormschema"
)

/*****************************************************************************************************************/

func main() {
	// Add the gorm schema to the loader here:
	stmts, err := gormschema.New("mysql").Load()

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, stmts)
}

/*****************************************************************************************************************/

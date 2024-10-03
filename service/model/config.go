/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package model

/*****************************************************************************************************************/

type Config struct {
	Debug bool   `default:"false"`
	Host  string `default:"0.0.0.0"`
	Port  int    `default:"50051"`
}

/*****************************************************************************************************************/
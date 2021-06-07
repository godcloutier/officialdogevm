/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


The tx creation was taken from https://www.thepolyglotdeveloper.com/2018/03/create-sign-bitcoin-transactions-golang/


*/
package cmd

import (
	"fmt"
	
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"	
)

// catimageCmd represents the config command
var catimageCmd = &cobra.Command{
	Use:   "catimage",
	Short: "creates a cat on the dogechain",
	Long: `The idea is to send a cat image over the dogechain
to have a proof of concept on a very basic level.

We want to bring a new protocol addon to the dogechain
which can coexist within transactions and does not need
any changes on the existing dogecoin-core.`,
	Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("CATIMAGE CALL")
	encodeImage(args[0]);
		
	},
}



func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func encodeImage(imagefile string) {
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile(imagefile)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"		
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)
}




func init() {
	rootCmd.AddCommand(catimageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catimageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catimageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

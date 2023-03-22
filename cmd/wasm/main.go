package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func createJSONFormatter() js.Func {
	formatter := js.FuncOf(formatJSONWrapper)
	return formatter
}

func formatJSONWrapper(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return buildErrorResponse("Incorrect number of arguments provided")
	}

	jsDocument, err := getDocument()
	if err != nil {
		return buildErrorResponse("Failed to obtain document object")
	}

	jsonOutputArea, err := getJSONOutputArea(jsDocument)
	if err != nil {
		return buildErrorResponse("Failed to access output text area")
	}

	jsonInput := args[0].String()
	fmt.Printf("Input: %s\n", jsonInput)
	formatted, err := formatJSON(jsonInput)
	if err != nil {
		errorMsg := fmt.Sprintf("Could not parse JSON. Error: %s\n", err)
		return buildErrorResponse(errorMsg)
	}

	jsonOutputArea.Set("value", formatted)
	return nil
}

func getDocument() (js.Value, error) {
	jsDocument := js.Global().Get("document")
	if !jsDocument.Truthy() {
		return js.Undefined(), fmt.Errorf("unable to get document object")
	}
	return jsDocument, nil
}

func getJSONOutputArea(jsDocument js.Value) (js.Value, error) {
	jsonOutputArea := jsDocument.Call("getElementById", "jsonOutput")
	if !jsonOutputArea.Truthy() {
		return js.Undefined(), fmt.Errorf("unable to get output text area")
	}
	return jsonOutputArea, nil
}

func buildErrorResponse(msg string) map[string]any {
	return map[string]any{
		"error": msg,
	}
}

func formatJSON(input string) (string, error) {
	var rawData any
	if err := json.Unmarshal([]byte(input), &rawData); err != nil {
		return "", err
	}
	formatted, err := json.MarshalIndent(rawData, "", "  ")
	if err != nil {
		return "", err
	}
	return string(formatted), nil
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("processJSON", createJSONFormatter())
	<-make(chan bool)
}

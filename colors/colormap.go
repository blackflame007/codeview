package colors

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var ColorMap = make(map[string]*color.Color)

var DefaultColors = map[string]string{
	"comment":                           "FgHiBlack",
	"number":                            "FgHiYellow",
	"type.builtin":                      "FgHiYellow",
	"label":                             "FgHiYellow",
	"warning":                           "FgHiYellow",
	"interpreted_string_literal":        "FgHiYellow",
	"var_declaration":                   "FgHiYellow",
	"var_spec":                          "FgHiYellow",
	"composite_literal":                 "FgHiYellow",
	"map_type":                          "FgHiYellow",
	"pointer_type":                      "FgHiYellow",
	"qualified_type":                    "FgHiYellow",
	"literal_value":                     "FgHiYellow",
	"keyed_element":                     "FgHiYellow",
	"literal_element":                   "FgHiYellow",
	"if_statement":                      "FgHiYellow",
	"binary_expression":                 "FgHiYellow",
	"short_var_declaration":             "FgHiYellow",
	"expression_list":                   "FgHiYellow",
	"index_expression":                  "FgHiYellow",
	"unary_expression":                  "FgHiYellow",
	"slice_type":                        "FgHiYellow",
	"slice_expression":                  "FgHiYellow",
	"for_clause":                        "FgHiYellow",
	"inc_statement":                     "FgHiYellow",
	"call_expression":                   "FgHiYellow",
	"selector_expression":               "FgHiYellow",
	"field_identifier":                  "FgHiYellow",
	"argument_list":                     "FgHiYellow",
	"\"":                                "FgHiYellow",
	"function":                          "FgHiCyan",
	"package":                           "FgHiCyan",
	"field":                             "FgHiCyan",
	"info":                              "FgHiCyan",
	"package_clause":                    "FgHiCyan",
	"composite_literal.type_identifier": "FgHiCyan",
	"type_identifier":                   "FgHiCyan",
	"call_expression.identifier":        "FgHiCyan",
	"selector_expression.identifier":    "FgHiCyan",
	"var_spec.type_identifier":          "FgHiCyan",
	"int_literal":                       "FgBlue",
	"nil":                               "FgBlue",
	"parameter":                         "FgBlue",
	"parameter_list":                    "FgBlue",
	"type":                              "FgHiGreen",
	"method":                            "FgHiGreen",
	"function_declaration":              "FgHiGreen",
	"function.call":                     "FgHiGreen",
	"method.call":                       "FgHiGreen",
	"function_declaration.identifier":   "FgHiGreen",
	"(":                                 "FgHiGreen",
	")":                                 "FgHiGreen",
	"keyword":                           "FgMagenta",
	"constant":                          "FgHiMagenta",
	"map":                               "FgHiMagenta",
	"import":                            "FgHiMagenta",
	"import_declaration":                "FgHiMagenta",
	"import_spec_list":                  "FgHiMagenta",
	"import_spec":                       "FgHiMagenta",
	":=":                                "FgHiMagenta",
	"var":                               "FgHiMagenta",
	"for":                               "FgHiMagenta",
	"if":                                "FgHiMagenta",
	"else":                              "FgHiMagenta",
	"func":                              "FgHiMagenta",
	"return_statement":                  "FgHiMagenta",
	"{":                                 "FgHiMagenta",
	"}":                                 "FgHiMagenta",
	"[":                                 "FgHiMagenta",
	"]":                                 "FgHiMagenta",
	"<":                                 "FgHiMagenta",
	">":                                 "FgHiMagenta",
	"=":                                 "FgHiMagenta",
	"==":                                "FgHiMagenta",
	"!=":                                "FgHiMagenta",
	"!":                                 "FgHiMagenta",
	"+":                                 "FgHiMagenta",
	"++":                                "FgHiMagenta",
	"-":                                 "FgHiMagenta",
	"*":                                 "FgHiMagenta",
	"return":                            "FgHiMagenta",
	"identifier":                        "FgWhite",
	".":                                 "FgWhite",
	",":                                 "FgWhite",
	";":                                 "FgWhite",
	":":                                 "FgWhite",
	"var_spec.identifier":               "FgWhite",
	"variable":                          "FgWhite",
	"operator":                          "FgWhite",
	"package_identifier":                "FgWhite",
	"parameter_declaration":             "FgWhite",
	"block":                             "FgWhite",
	"debug":                             "FgWhite",
	"source_file":                       "FgWhite",
	"escape_sequence":                   "FgRed",
	"string":                            "FgRed",
	"escape.character":                  "FgRed",
	"error":                             "FgRed",

	// ... add more as needed
}

func InitColors() {
	// Set defaults in Viper
	for nodeType, colorValue := range DefaultColors {
		viper.SetDefault("colors."+nodeType, colorValue)
	}

	// Initialize ColorMap from Viper
	colorConfig := viper.GetStringMapString("colors")
	for nodeType, colorName := range colorConfig {
		ColorMap[nodeType] = parseColor(colorName)
	}
}

func parseColor(colorName string) *color.Color {
	switch colorName {
	case "FgBlack":
		return color.New(color.FgBlack)
	case "FgRed":
		return color.New(color.FgRed)
	case "FgGreen":
		return color.New(color.FgGreen)
	case "FgYellow":
		return color.New(color.FgYellow)
	case "FgBlue":
		return color.New(color.FgBlue)
	case "FgMagenta":
		return color.New(color.FgMagenta)
	case "FgCyan":
		return color.New(color.FgCyan)
	case "FgWhite":
		return color.New(color.FgWhite)
	case "FgHiBlack":
		return color.New(color.FgHiBlack)
	case "FgHiRed":
		return color.New(color.FgHiRed)
	case "FgHiGreen":
		return color.New(color.FgHiGreen)
	case "FgHiYellow":
		return color.New(color.FgHiYellow)
	case "FgHiBlue":
		return color.New(color.FgHiBlue)
	case "FgHiMagenta":
		return color.New(color.FgHiMagenta)
	case "FgHiCyan":
		return color.New(color.FgHiCyan)
	case "FgHiWhite":
		return color.New(color.FgHiWhite)
	default:
		return color.New(color.FgWhite)
	}
}

// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr" /* copybara-comment: antlr */
)

// LexerListener listens for and formats SyntaxErrors in a way suitable for
// the antlr Lexer.
type LexerListener struct {
	*antlr.DefaultErrorListener
	Code string
}

// SyntaxError is called when the parser encounters a syntax error.
func (l *LexerListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	// We can possibly use offendingSymbol, but this is unreliably populated.
	panic(NewTranspilationError(line, column, fmt.Errorf("%q is not expected here. %s", nextToken(l.Code, line, column), msg)))
}

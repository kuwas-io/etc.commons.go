/** @imports */
package tools
import ( "io/ioutil" )

/** @parameters */
var title = "Tools"

/** @exports */
func Default( s string ) ( string , error ) {
  var e error
  var o [] byte
  if o , e = ioutil.ReadFile( s ) ; e != nil { return "" , e }
  return string( o ) , nil
}

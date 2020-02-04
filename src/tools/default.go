/** @imports */
package tools
import ( "io/ioutil" )

/** @parameters */
var title = "Tools"

/** @exports */
func Default( o string ) ( r string , e error ) {
  var b [] byte
  if b , e = ioutil.ReadFile( o ) ; e != nil { return }
  r = string( b )
  return
}

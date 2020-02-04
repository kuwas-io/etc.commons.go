/** @imports */
package tools
import ( "strings" )

/** @parameters */
/** @exports */
func Title( o [] string , t bool ) ( r string ) {
  var d = " » "
  r = strings.Join( o , d )
  if t { r += d }
  return
}

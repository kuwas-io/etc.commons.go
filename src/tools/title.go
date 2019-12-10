/** @imports */
package tools
import ( "strings" )

/** @parameters */
/** @exports */
func Title( s [] string , t bool ) string {
  var d = " »"
  var o = strings.Join( s , d )
  if t { o += d }
  return o
}

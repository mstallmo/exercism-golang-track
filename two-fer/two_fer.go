//twofer package contains a function ShareWith that takes a string returns a string that inserts the argument into the
// sentence "One for <name>, one for me.
package twofer

//Share with takes a string as an argument and returns "One for <name>, one for me" if a string is passed in. If the
//argument is left empty the function will return "One for you, one for me"
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}

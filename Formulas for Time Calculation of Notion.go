if(
	empty(prop("State")) == true, 
	"Please Set the State", 
	if(
		empty(prop("Date")) == true, 
		"Please Set the Start Date", 
		if(
			prop("State") == "Wish List", 
			if(
				largerEq(dateBetween(prop("Date"), now(), "days"), 1), 
				if(
					larger(dateBetween(prop("Date"), now(), "days"), 1), 
					format(dateBetween(prop("Date"), now(), "days")) + " days to go", 
					format(dateBetween(prop("Date"), now(), "days")) + " day to go"
				), 
				"You Should Start Right Now"
			), 	
			if(
				prop("State") == "In Progress", 
				if(
					largerEq(dateBetween(now(), prop("Date"), "days"), 1), 
					if(
						larger(dateBetween(now(), prop("Date"), "days"), 1), 
						format(dateBetween(now(), prop("Date"), "days")) + " days have passed", 
						format(dateBetween(now(), prop("Date"), "days")) + " day has passed"
					),
					if(
						smaller(dateBetween(prop("Date"), now(), "days"), 0), 
						"You Should Start Right Now", 
						"A Journey of A Thousand Miles Begins with A Single Step"
					)
				), 
				if(
					prop("State") == "Archived", 
					if(
						largerEq(dateBetween(end(prop("Date")), prop("Date"), "days"), 1), 
						if(
							larger(dateBetween(end(prop("Date")), prop("Date"), "days"), 1), 
							format(dateBetween(end(prop("Date")), prop("Date"), "days")) + " days were used", 
							format(dateBetween(end(prop("Date")), prop("Date"), "days")) + " day was used"
						), 
						if(
							smaller(dateBetween(prop("Date"), end(prop("Date")), "days"), 0), 
							"Error", 
							"Read Efficiently, Read Effectively" 
						)
					), 
					"Error"	
				)		
			)
		)
	)
)






if(empty(prop("State")) == true, "Please Set the State", if(empty(prop("Date")) == true, "Please Set the Start Date", if(prop("State") == "Wish List", if(largerEq(dateBetween(prop("Date"), now(), "days"), 1), if(larger(dateBetween(prop("Date"), now(), "days"), 1), format(dateBetween(prop("Date"), now(), "days")) + " days to go", format(dateBetween(prop("Date"), now(), "days")) + " day to go"), "You Should Start Right Now"), if(prop("State") == "In Progress", if(largerEq(dateBetween(now(), prop("Date"), "days"), 1), if(larger(dateBetween(now(), prop("Date"), "days"), 1), format(dateBetween(now(), prop("Date"), "days")) + " days have passed", format(dateBetween(now(), prop("Date"), "days")) + " day has passed"), if(smaller(dateBetween(prop("Date"), now(), "days"), 0), "You Should Start Right Now", "A Journey of A Thousand Miles Begins with A Single Step")), if(prop("State") == "Archived", if(largerEq(dateBetween(end(prop("Date")), prop("Date"), "days"), 1), if(larger(dateBetween(end(prop("Date")), prop("Date"), "days"), 1), format(dateBetween(end(prop("Date")), prop("Date"), "days")) + " days were used", format(dateBetween(end(prop("Date")), prop("Date"), "days")) + " day was used" ),if(smaller(dateBetween(prop("Date"), end(prop("Date")), "days"), 0), "Error", "Read Efficiently, Read Effectively")), "Error")))))





// and ???
// or ???
// not ???
// use the format() to transform number to text
// + "text"

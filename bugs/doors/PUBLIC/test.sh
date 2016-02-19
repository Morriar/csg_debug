#!/bin/sh

for i in $(ls $1); do
	case $i in
		*a*)
			# Each line of the file contains a name, its
			# allegiance index and its personal accepted
			# allegiance ranges.
			#
			# Reject those with out of bound allegiance index.

			cat $i | while read name idx min max; do
				if [ $idx -ge $min && $idx -le $min ]; then
					echo ACCEPT $name $i
				else
					echo REJECT $name $i
				fi
			done
			;;

		*b*)
			# The file contains a citizen information with its
			# name, its address and its occupation.
			#
			# Accepts only citizens that are more equal than other.
			
			if [ grep MINISTER $i ]; then
				echo ACCEPT MINISTER $i
			elif [ grep MILITIA $i ]; then
				echo ACCEPT MILITIA $i
			else
				echo REJECT $i
			fi;;

		*c*)
			# The file contains the citizenship log of the
			# citizenship level-related orders.
			#
			# The level can be increased by a number (promotion),
			# or decreased by a factor (demotion).
			#
			# Accept only citizen with a citizenship level over 9000.

			cat $i | while read order value; do
				if [ $order == PROMOTE ]; then
					level=$((level + value))
				fi
				if [ $order == DEMOTE ]; then
					level=$((level / value))
				fi
			done
			if [ $level -gt 9000 ]; then
				echo ACCEPT LEVEL $level $i
			else
				echo REJECT LEVEL $level $i
			fi
			;;

		*d*)
			# Datum contains a list of citizens.
			#
			# Accept only the 10 more frequent citizens (output the most frequent last).

			# 1. Down-case the file
			# 2. Count number of occurrences
			# 3. Sort by number of occurrences
			# 4. Keep 10 best entries
			# 5. Remove number of occurrences
			for n in `tr [A-z] [a-z] $i | uniq -c | sort | head -n 10 | cut -c 9`; do
				echo ACCEPT $n $i
			done
			;;
		*)
			echo UNEXPECTED DATUM $i. FALLBACK TO KILLING PEOPLE.
			exit 666
	esac
done
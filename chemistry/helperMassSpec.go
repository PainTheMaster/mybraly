package chemistry

import (
	"PainTheMaster/mybraly/order"
)

const (
	methodBoudary = 16
)

//isotopeMerge merges two isotoep patterns. If match in mass is found the intensity of the element in addedFrom is added to the corresponding element in  *ptrAddedTo
func isotopeMarge(ptrAddedTo *Isotopes, addedFrom Isotopes) {
	addedTo := *ptrAddedTo

	for _, isotopeAdded := range addedFrom {
		boolFound, idxFound := isotopemassBinarySearch(addedTo, isotopeAdded)
		if boolFound {
			addedTo[idxFound].Abundance += isotopeAdded.Abundance
		} else {
			addedTo = append(addedTo, isotopeAdded)
		}
	}
	*ptrAddedTo = addedTo
}

func isotopemassBinarySearch(pool Isotopes, key Isotope) (boolFound bool, idxFound int) {
	order.QuickSort(pool)
	ini := 0
	end := len(pool) - 1
	middle := ini + (end-ini)/2

	boolFound = false
	if end-ini+1 <= methodBoudary {
		boolFound, idxFound = linearSearch(pool, key)
	} else {
		for {
			if end-ini+1 <= methodBoudary {
				boolFound, idxFound = linearSearch(pool, key)
				break
			}
			if key.Mass < pool[middle].Mass {
				end = middle
			} else if key.Mass > pool[middle].Mass {
				ini = middle
			} else {
				boolFound = true
				idxFound = middle
			}
		}
	}
	return
}

func linearSearch(isotopePattern Isotopes, singleIsotope Isotope) (boolFound bool, idxFound int) {
	boolFound = false
	for idx, isotope := range isotopePattern {
		if isotope.Mass == singleIsotope.Mass {
			boolFound = true
			idxFound = idx
			break
		}
	}
	return
}

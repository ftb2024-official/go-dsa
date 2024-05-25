package dynamicarray

import "errors"

var defaultCapacity = 10

// "representing DynamicArray
type DynamicArray struct {
	Size        int
	Capacity    int
	ElementData []any
}

// CheckRangeFromIndex checks the range from index
func (da *DynamicArray) CheckRangeFromIndex(index int) error {
	if index >= da.Size || index < 0 {
		return errors.New("index out of range")
	}

	return nil
}

// NewCapacity increments the capacity of dynamic array
func (da *DynamicArray) NewCapacity() {
	if da.Capacity == 0 {
		da.Capacity = defaultCapacity
	} else {
		da.Capacity <<= 1 // da.Capacity = da.Capacity << 1
	}

	newElementData := make([]any, da.Capacity)
	copy(newElementData, da.ElementData)
	da.ElementData = newElementData
}

// IsEmpty checks if the dynamic array is empty
func (da *DynamicArray) IsEmpty() bool {
	return da.Size == 0
}

// Put updates the value in given index to the given value
func (da *DynamicArray) Put(index int, element any) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	da.ElementData[index] = element

	return nil
}

// Add adds element to the dynamic array
func (da *DynamicArray) Add(element any) {
	if da.Size == da.Capacity {
		da.NewCapacity()
	}

	da.ElementData[da.Size] = element
	da.Size++
}

// Remove removes element from the dynamic array
func (da *DynamicArray) Remove(index int) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	copy(da.ElementData[index:], da.ElementData[index+1:da.Size])
	da.ElementData[da.Size-1] = nil

	da.Size--

	return nil
}

// Get returns the specific element from dynamic array
func (da *DynamicArray) Get(index int) (any, error) {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return nil, err
	}

	return da.ElementData[index], nil
}

// GetData returns all elements from dynamic array
func (da *DynamicArray) GetData() []any {
	return da.ElementData[:da.Size]
}

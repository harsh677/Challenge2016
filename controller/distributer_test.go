package controller

import (
	"example/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckByPermissionTypeInc(t *testing.T) {
	// Test case 1: Check inclusion for a valid country, should return true
	parentPermission1 := models.Permission{
		IncludeMap: map[string]models.PermissionType{"India": models.CountryType},
		ExcludeMap: map[string]models.PermissionType{},
	}

	countryStateMap1 := make(models.CountryMap)
	countryStateMap1["India"] = make(models.StateMap)

	result1, err1 := checkByPermissionTypeInc("India", parentPermission1, countryStateMap1)
	assert.True(t, result1)
	assert.Nil(t, err1)

	// Test case 2: Check inclusion for an invalid country, should return false with an error
	parentPermission2 := models.Permission{
		IncludeMap: map[string]models.PermissionType{"USA": models.CountryType},
		ExcludeMap: map[string]models.PermissionType{},
	}

	countryStateMap2 := make(models.CountryMap)
	countryStateMap2["India"] = make(models.StateMap)

	result2, err2 := checkByPermissionTypeInc("India", parentPermission2, countryStateMap2)
	assert.False(t, result2)
	assert.NotNil(t, err2) // Fix: An error is expected here
	assert.EqualError(t, err2, "invalid input country India")
}

func TestCheckByPermissionTypeExc(t *testing.T) {
	// Test case 1: Check exclusion for a valid country, should return true
	parentPermission1 := models.Permission{
		IncludeMap: map[string]models.PermissionType{},
		ExcludeMap: map[string]models.PermissionType{"India": models.CountryType},
	}

	countryStateMap1 := make(models.CountryMap)
	countryStateMap1["India"] = make(models.StateMap)

	result1, err1 := checkByPermissionTypeExc("India", parentPermission1, countryStateMap1)
	assert.True(t, result1)
	assert.Nil(t, err1)

	// Test case 2: Check exclusion for an invalid country, should return false with an error
	parentPermission2 := models.Permission{
		IncludeMap: map[string]models.PermissionType{},
		ExcludeMap: map[string]models.PermissionType{"USA": models.CountryType},
	}

	countryStateMap2 := make(models.CountryMap)
	countryStateMap2["India"] = make(models.StateMap)

	result2, err2 := checkByPermissionTypeExc("India", parentPermission2, countryStateMap2)
	assert.False(t, result2)
	assert.EqualError(t, err2, "invalid input country India")
}

func TestCheckPermission(t *testing.T) {
	// Test case 1: Check permission with valid inclusion, should return true
	child1 := models.Distributer{
		InputPermission: "India",
		AuthType:        models.Include,
	}

	parentPermission1 := models.Permission{
		IncludeMap: map[string]models.PermissionType{"India": models.CountryType},
		ExcludeMap: map[string]models.PermissionType{},
	}

	countryStateMap1 := make(models.CountryMap)
	countryStateMap1["India"] = make(models.StateMap)

	result1, err1 := checkPermission(child1, parentPermission1, countryStateMap1)
	assert.True(t, result1)
	assert.Nil(t, err1)

	// Test case 2: Check permission with invalid exclusion, should return false with an error
	child2 := models.Distributer{
		InputPermission: "India",
		AuthType:        models.Exclude,
	}

	parentPermission2 := models.Permission{
		IncludeMap: map[string]models.PermissionType{},
		ExcludeMap: map[string]models.PermissionType{"India": models.CountryType},
	}

	countryStateMap2 := make(models.CountryMap)
	countryStateMap2["India"] = make(models.StateMap)

	result2, err2 := checkPermission(child2, parentPermission2, countryStateMap2)
	assert.False(t, result2)
	assert.EqualError(t, err2, "Parent distributer dont have access to grant permission- India")
}

func TestAddAnyDist(t *testing.T) {
	// Test case 1: Add distributor with valid inclusion, should return nil error
	distributer1 := models.Distributer{
		Name:            "Vyara",
		InputPermission: "India",
		AuthType:        models.Include,
	}

	distributerMap1 := make(models.DistributerMap)

	err1 := addAnyDist(distributer1, distributerMap1, 0, nil)
	assert.Nil(t, err1)

	// Test case 2: Add distributor with existing inclusion, should return error
	distributer2 := models.Distributer{
		Name:            "Vyara",
		InputPermission: "India",
		AuthType:        models.Include,
	}

	distributerMap2 := make(models.DistributerMap)
	distributerMap2["Vyara"] = models.Permission{
		IncludeMap: map[string]models.PermissionType{"India": models.CountryType},
		ExcludeMap: map[string]models.PermissionType{},
	}

	err2 := addAnyDist(distributer2, distributerMap2, 0, nil)
	assert.EqualError(t, err2, "permission Exist.. in include - India")
}

func TestGetPermissionType(t *testing.T) {
	// Test case 1: Get permission type for a country, should return CountryType
	pArr1, pType1 := getPermissionType("India")
	assert.Equal(t, []string{"India"}, pArr1)
	assert.Equal(t, models.CountryType, pType1)

	// Test case 2: Get permission type for a state, should return StateType
	pArr2, pType2 := getPermissionType("Gujarat-India")
	assert.Equal(t, []string{"Gujarat", "India"}, pArr2)
	assert.Equal(t, models.StateType, pType2)

	// Test case 3: Get permission type for a city, should return CityType
	pArr3, pType3 := getPermissionType("Vyara-Gujarat-India")
	assert.Equal(t, []string{"Vyara", "Gujarat", "India"}, pArr3)
	assert.Equal(t, models.CityType, pType3)
}

func TestUpperCaseNoSpace(t *testing.T) {
	// Test case 1: Convert to uppercase without spaces
	result1 := UpperCaseNoSpace("test string")
	assert.Equal(t, "TESTSTRING", result1)

	// Test case 2: Convert to uppercase without spaces for an already uppercase string
	result2 := UpperCaseNoSpace("UPPERCASE")
	assert.Equal(t, "UPPERCASE", result2)
}

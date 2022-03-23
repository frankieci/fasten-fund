package stream

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Created by frankieci on 2022/3/16 11:53 pm
 */

type TestAlarm struct {
	AlarmCount   int64
	ResolveCount int64
	WarnType     string
	Period       string
}

func (a TestAlarm) Identify() interface{} {
	return a.AlarmCount / 10
}

func TestJust(t *testing.T) {
	alarms := []Element{
		TestAlarm{AlarmCount: 1, ResolveCount: 0},
		TestAlarm{AlarmCount: 10, ResolveCount: 5},
		TestAlarm{AlarmCount: 20, ResolveCount: 10},
	}
	result, err := Just(alarms...).Reduce(func(pipe <-chan Element) (Element, error) {
		result := TestAlarm{}
		for e := range pipe {
			alarm := e.(TestAlarm)
			result.AlarmCount = result.AlarmCount + alarm.AlarmCount
			result.ResolveCount = result.ResolveCount + alarm.ResolveCount
		}
		return result, nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	expected := TestAlarm{AlarmCount: 31, ResolveCount: 15}
	assert.Equal(t, expected, result)
}

func TestGroup(t *testing.T) {
	alarms := []Element{
		TestAlarm{AlarmCount: 1, ResolveCount: 0, Period: "2020-02-01"},
		TestAlarm{AlarmCount: 10, ResolveCount: 5, Period: "2020-02-02"},
		TestAlarm{AlarmCount: 20, ResolveCount: 10, Period: "2020-02-01"},
	}

	result, err := Just(alarms...).Group(func(e Element) Key {
		alarm := e.(TestAlarm)
		return alarm.Period
	}).Reduce(func(pipe <-chan Element) (Element, error) {
		result := Elements{}
		for es := range pipe {
			groupOne := TestAlarm{}
			elements := es.(Elements)
			for i, e := range elements {
				alarm := e.(TestAlarm)
				if i == 0 {
					groupOne.Period = alarm.Period
				}
				groupOne.AlarmCount = groupOne.AlarmCount + alarm.AlarmCount
			}
			result = append(result, groupOne)
		}
		return result, nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	expected := Elements{
		TestAlarm{AlarmCount: 21, ResolveCount: 0, Period: "2020-02-01"},
		TestAlarm{AlarmCount: 10, ResolveCount: 0, Period: "2020-02-02"},
	}

	assert.Equal(t, expected, result)
}

func TestSort(t *testing.T) {
	alarms := []Element{
		TestAlarm{AlarmCount: 1, ResolveCount: 0},
		TestAlarm{AlarmCount: 20, ResolveCount: 5},
		TestAlarm{AlarmCount: 10, ResolveCount: 10},
	}

	sorted := Elements{}
	Just(alarms...).Sort(func(a, b Element) bool {
		return a.(TestAlarm).AlarmCount < b.(TestAlarm).AlarmCount
	}).ForEach(func(e Element) {
		sorted = append(sorted, e)
	})

	expected := Elements{
		TestAlarm{AlarmCount: 1, ResolveCount: 0},
		TestAlarm{AlarmCount: 10, ResolveCount: 10},
		TestAlarm{AlarmCount: 20, ResolveCount: 5},
	}
	assert.Equal(t, expected, sorted)
}

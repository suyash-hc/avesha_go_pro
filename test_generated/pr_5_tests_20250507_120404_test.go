package test_generated

import (\n\t"testing"\n\t"test-avesha-agent/math"\n)

func TestAdd(t *testing.T) {\n\tcases := []struct {\n\t\ta, b, expected int\n\t}{\n\t\t{1, 2, 3},\n\t\t{0, 0, 0},\n\t\t{-1, 1, 0},\n\t}\n\n\tfor _, c := range cases {\n\t\tresult := math.Add(c.a, c.b)\n\t\tif result != c.expected {\n\t\t\tt.Errorf("Add(%d, %d) = %d, expected %d", c.a, c.b, result, c.expected)\n\t\t}\n\t}\n}

func TestSubtract(t *testing.T) {\n\tcases := []struct {\n\t\ta, b, expected int\n\t}{\n\t\t{5, 2, 3},\n\t\t{0, 0, 0},\n\t\t{10, 5, 5},\n\t}\n\n\tfor _, c := range cases {\n\t\tresult := math.Subtract(c.a, c.b)\n\t\tif result != c.expected {\n\t\t\tt.Errorf("Subtract(%d, %d) = %d, expected %d", c.a, c.b, result, c.expected)\n\t\t}\n\t}\n}




# Generated tests for PR 1
# Generated on: 2025-05-05 18:12:45


# Test: test_generated
The generated tests for the git diff are as follows:

For README.md:
```python
def test_readme_content():
    assert 'Test Avesha Agent' in readme_content
    assert 'A basic Go project template.' in readme_content
    assert 'Getting Started' in readme_content
    assert 'Go 1.21 or later' in readme_content
    assert 'Installation' in readme_content
    assert 'Running the Application' in readme_content
    assert 'Project Structure' in readme_content
    assert 'License' in readme_content
```

For go.mod:
```python
def test_go_mod_version():
    assert 'go 1.21' in go_mod_content
```

For main.go:
```python
def test_main_functionality():
    assert math.Add(10, 5) == 15
    assert math.Subtract(10, 5) == 5
    assert math.Multiply(10, 5) == 50
    assert math.Divide(10, 5) == 2
```

For operations.go:
```python
def test_addition():
    assert math.Add(3, 5) == 8

def test_subtraction():
    assert math.Subtract(10, 3) == 7

def test_multiplication():
    assert math.Multiply(4, 6) == 24

def test_division():
    assert math.Divide(10, 2) == 5
    with pytest.raises(ZeroDivisionError):
        math.Divide(10, 0)
```

package programs

import (
	"fmt"
	"terminal-emulator/vfs"
	"testing"
)

var testVFS string = `<node name="" dir="true" modified="2025-09-17T15:24:55.4406365+03:00">
	<node name="home" dir="true" modified="2025-09-17T15:24:55.4406365+03:00">
		<node name="user" dir="true" modified="2025-09-17T15:24:55.4406365+03:00">
		<node name="image.jpg" dir="false" modified="2025-09-17T15:24:55.4406365+03:00">
			<content>image12345</content>
		</node>
		<node name="empty" dir="true" modified="2025-09-17T15:24:55.4406365+03:00"></node>
		<node name="empty.txt" dir="false" modified="2025-09-18T13:32:11.0540706+03:00"></node>
		<node name="ff.txt" dir="false" modified="2025-09-20T13:24:27.524069+03:00"></node>
		</node>
		<node name="touch" dir="false" modified="2025-09-19T12:52:13.9736937+03:00"></node>
	</node>
	<node name="etc" dir="true" modified="2025-09-17T15:24:55.4406365+03:00">
		<node name="config.conf" dir="false" modified="2025-09-17T15:24:55.4406365+03:00">
		<content>configuration</content>
		</node>
	</node>
	<node name="document.txt" dir="false" modified="2025-09-20T13:24:15.4254299+03:00">
		<content>doc1&#xA;doc2&#xA;doc3&#xA;doc4&#xA;doc5&#xA;doc6&#xA;doc7&#xA;doc8&#xA;doc9&#xA;doc10&#xA;doc11</content>
	</node>
</node>`

func setupVFS() {
	fs := vfs.LoadFromString([]byte(testVFS))
	vfs.SetupExplorer(fs)
}

func TestProgramsOutput(t *testing.T) {
	tests := []struct {
		name     string
		program  func()
		input    []string
		expected []string
		isError  bool
	}{
		{
			name:     "LsTest",
			program:  program(Ls),
			input:    nil,
			expected: []string{"home", "etc", "document.txt"},
			isError:  false,
		},
		{
			name:     "LsWrongArgTest",
			program:  program(Ls),
			input:    []string{"wrong"},
			expected: nil,
			isError:  true,
		},
		{
			name:     "CatTest",
			program:  program(Cat),
			input:    []string{"document.txt"},
			expected: []string{"doc1", "doc2", "doc3", "doc4", "doc5", "doc6", "doc7", "doc8", "doc9", "doc10", "doc11"},
			isError:  false,
		},
		{
			name:     "CatWrongTest",
			program:  program(Cat),
			input:    []string{"1"},
			expected: nil,
			isError:  true,
		},
		{
			name:     "TailTest",
			program:  program(Tail),
			input:    []string{"document.txt"},
			expected: []string{"doc2", "doc3", "doc4", "doc5", "doc6", "doc7", "doc8", "doc9", "doc10", "doc11"},
			isError:  false,
		},
		{
			name:     "DuTest",
			program:  program(Du),
			input:    nil,
			expected: []string{"    10    home", "    13    etc", "    56    document.txt"},
			isError:  false,
		},
	}

	setupVFS()
	output := make([]string, 0)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Execute(test.program, test.input, func(i any) {
				output = append(output, fmt.Sprint(i))
			})

			if !test.isError && err != nil {
				t.Errorf("\nunexpected error: %s", err)
			}
			if test.isError && err == nil {
				t.Error("\nexpected error, got nil")
			}

			if len(output) != len(test.expected) {
				t.Errorf("\nexpected: %d lines, got: %d lines", len(test.expected), len(output))
			} else {
				for i := range output {
					if output[i] != test.expected[i] {
						t.Errorf("\nexpected: %s, got: %s", test.expected[i], output[i])
					}
				}
			}
		})
		output = make([]string, 0)
	}
}

package programs

import (
	"fmt"
	"terminal-emulator/vfs"
	"testing"
)

var testVFS string = `<node name="" dir="true" modified="2025-09-17T15:24:55.4406365+03:00">
	<node name="home" dir="true" modified="2025-09-17T15:24:55.4406365+03:00">
		<node name="user" dir="true" modified="2025-09-17T15:24:55.4406365+03:00">
		<node name="document.txt" dir="false" modified="2025-09-20T13:24:15.4254299+03:00">
			<content>doc1&#xA;doc2&#xA;doc3&#xA;doc4&#xA;doc5&#xA;doc6&#xA;doc7&#xA;doc8&#xA;doc9&#xA;doc10&#xA;doc11</content>
		</node>
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
</node>`

func setupVFS() {
	fs := vfs.LoadFromString([]byte(testVFS))
	vfs.SetupExplorer(fs)
}

func TestLs(t *testing.T) {
	out := make([]string, 0, 2)
	correct := []string{"home", "etc"}
	ls := program(Ls)

	setupVFS()

	err := Execute(ls, nil, func(i any) {
		out = append(out, fmt.Sprint(i))
	})

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	for i := range out {
		if out[i] != correct[i] {
			t.Errorf("\nexpected: %s\ngot: %s", correct[i], out[i])
		}
	}
}

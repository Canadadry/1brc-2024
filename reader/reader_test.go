package reader

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	for version, impl := range rsFunctions {
		t.Run(version, func(t *testing.T) {
			tests := map[string]struct {
				input          string
				expectedOutput string
			}{
				"single station": {
					input:          "StationA;20.5\nStationA;22.5\n",
					expectedOutput: "{StationA=20.5/21.5/22.5}\n",
				},
				"multiple stations": {
					input:          "StationA;20.5\nStationB;25.5\nStationA;22.5\nStationB;27.5\n",
					expectedOutput: "{StationA=20.5/21.5/22.5, StationB=25.5/26.5/27.5}\n",
				},
				"invalid lines ignored": {
					input:          "StationA;20.5\nInvalidLine\nStationA;22.5\n",
					expectedOutput: "{StationA=20.5/21.5/22.5}\n",
				},
				"alphabetical order": {
					input:          "StationC;5.0\nStationA;3.5\nStationB;4.2\nStationA;2.5\nStationC;5.2\n",
					expectedOutput: "{StationA=2.5/3.0/3.5, StationB=4.2/4.2/4.2, StationC=5.0/5.1/5.2}\n",
				},
				"one decimal digit": {
					input:          "Station1;12.34\n",
					expectedOutput: "{Station1=12.3/12.3/12.3}\n",
				},
				"positive and negative temperatures": {
					input:          "Station1;-5.0\nStation2;3.5\nStation1;2.0\nStation2;-2.5\nStation3;0.0\n",
					expectedOutput: "{Station1=-5.0/-1.5/2.0, Station2=-2.5/0.5/3.5, Station3=0.0/0.0/0.0}\n",
				},
			}
			for name, tc := range tests {
				t.Run(name, func(t *testing.T) {
					input := strings.NewReader(tc.input)
					var output bytes.Buffer

					err := impl(input, &output)
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}

					if gotOutput := output.String(); gotOutput != tc.expectedOutput {
						t.Errorf("expected output %q, got %q", tc.expectedOutput, gotOutput)
					}
				})
			}
		})
	}
}

func BenchmarkRead(b *testing.B) {
	inputData := "StationA;1.1\nStationB;2.2\nStationC;3.3\nStationA;4.4\nStationB;5.5\n"
	for i := 0; i < 1000; i++ {
		inputData += "StationA;1.1\nStationB;2.2\nStationC;3.3\nStationA;4.4\nStationB;5.5\n"
	}

	for name, impl := range rsFunctions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r := strings.NewReader(inputData)
				w := ioutil.Discard
				err := impl(r, w)
				if err != nil {
					b.Fatal("failed:", err)
				}
			}
		})
	}
}

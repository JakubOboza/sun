package vcgencmd

import (
  "os/exec"
  "strings"
  "regexp"
  "bytes"
  "strconv"
)

func RunCommand(command string) (string, error) {
  cmd := exec.Command("/opt/vc/bin/vcgencmd", command)
  cmd.Stdin = strings.NewReader("cpu temperature value")
  var out bytes.Buffer
  cmd.Stdout = &out

  err := cmd.Run()
  if err != nil {
    return "", err
  }

  return out.String(), nil
}

func MeasureTempString() (string, error) {
  return RunCommand("measure_temp")
}

func MeasureTemp() (float64, error) {
 raw_value, err :=  MeasureTempString()

 if err != nil {
   return 0.0, err
 }
 temp_regex, _ := regexp.Compile("^temp=(.+)\\'C$") // temp=46.2'C
 temp_value := temp_regex.FindStringSubmatch(strings.TrimSpace(raw_value))[1]

 return strconv.ParseFloat(temp_value, 64)
}

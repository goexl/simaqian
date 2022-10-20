package simaqian

var (
	_ = Output

	_ option = (*optionOutput)(nil)
)

type optionOutput struct {
	output writer
}

// Output 输出流
func Output(output writer) *optionOutput {
	return &optionOutput{
		output: output,
	}
}

func (o *optionOutput) apply(options *options) {
	options.outputs = append(options.outputs, o.output)
}

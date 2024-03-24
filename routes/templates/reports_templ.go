// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.646
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ReportsView() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div data-bs-theme=\"dark\"><h3 style=\"color: #FFFFFF;\">PICT Library - Generate Report</h3><div style=\"height: 20px;\"></div><form id=\"submitForm\"><div class=\"input-group mb-3\"><input type=\"text\" id=\"location\" class=\"form-control\" placeholder=\"Location\" aria-label=\"Location\"></div><div class=\"input-group mb-3\"><input type=\"date\" id=\"from-date\" class=\"form-control\" placeholder=\"From Date\" aria-label=\"From Date\"></div><div class=\"input-group mb-3\"><input type=\"date\" id=\"to-date\" class=\"form-control\" placeholder=\"To Date\" aria-label=\"To Date\"></div><button class=\"btn btn-primary\" type=\"submit\" id=\"enter-btn\">Generate</button><div style=\"height: 10px;\"></div><a type=\"button\" href=\"/\" class=\"btn btn-secondary\">Home</a><div id=\"dlDiv\"></div></form></div><script>\n      let dlDiv = document.getElementById(\"dlDiv\")\n      document.getElementById(\"submitForm\").addEventListener(\"submit\", (e) => {\n        e.preventDefault()\n        let location = document.getElementById(\"location\").value\n        let fromDate = new Date(document.getElementById(\"from-date\").value)\n        let toDate = new Date(document.getElementById(\"to-date\").value)\n        toDate.setDate(toDate.getDate() + 1)\n        axios.post(\"/report\", {\n          location,\n          start_time: fromDate.toISOString(),\n          stop_time: toDate.toISOString(),\n        }, { responseType: \"blob\" }).then(res => {\n          const url = window.URL.createObjectURL(new Blob([res.data]))\n          let a = document.createElement(\"a\")\n          a.href = url\n          a.download = \"Report.xlsx\"\n          dlDiv.appendChild(a)\n          a.click()\n          window.URL.revokeObjectURL(url);\n        }).catch(err => {\n          console.error(err)\n        })\n      })\n    </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Layout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
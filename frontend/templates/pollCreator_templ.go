// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreatorChartView(title string, answers []string, pollResults []int, pollId string, env string) templ.Component {
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
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, loadChart(answers, pollResults, pollId))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body onload=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.ComponentScript = loadChart(answers, pollResults, pollId)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><h1 class=\"flex justify-center text-center px-5 pt-14 pb-10 min-[601px]:pt-28 min-[601px]:pb-20 text-3xl font-extrabold underline\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string = title
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><div class=\"bg-[#FCFDFC] p-4\"><div class=\"flex justify-center items-center relative m-auto h-[50vh] max-w-[1200px] px-2 min-[601px]:px-4 shadow-lg shadow-slate-200 rounded\"><canvas id=\"myChart1\"></canvas></div></div><div class=\"px-4\"><div class=\"flex flex-col justify-center items-center text-center mt-20 mb-5\"><h1 class=\"text-2xl\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := `Want to easily share this poll with others? Copy the link below!`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></div><div class=\"m-auto max-w-[800px] bg-slate-50 border-2 border-[#809D80] rounded-sm px-4\"><div class=\"flex justify-between items-center py-1\"><code>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string = getUrlToCopy(pollId, env)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</code><div class=\"flex items-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, copyUrl(getUrlToCopy(pollId, env)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button id=\"copyURLButton\" class=\"w-[120px] py-1 px-2 bg-[#809D80] text-zinc-50 duration-300 cursor-pointer border border-slate-950 hover:bg-[#5c735c] hover:text-white rounded-sm\" onclick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 templ.ComponentScript = copyUrl(getUrlToCopy(pollId, env))
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var6.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" role=\"button\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var7 := `Copy`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var7)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <img src=\"../assets/images/copy.png\" width=\"20\" height=\"20\" class=\" mt-1 float-right\"></button></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = openEventConnection(pollId).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!--")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var8 := ` icons from <a href="https://www.flaticon.com/free-icons/copy" title="copy icons">Copy icons created by Freepik - Flaticon</a> `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var8)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("--></body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func PollCreator(title string, answers []string, pollResults []int, pollId string, env string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var10 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = CreatorChartView(title, answers, pollResults, pollId, env).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Head("GoPoll | Your Poll").Render(templ.WithChildren(ctx, templ_7745c5c3_Var10), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func openEventConnection(pollId string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_openEventConnection_6e34`,
		Function: `function __templ_openEventConnection_6e34(pollId){const eventSource = new EventSource("/updateChart/" + pollId + "?stream="+pollId);

    eventSource.onopen = function() {
        console.log("event openned")
    };

    eventSource.onmessage = (event) => {
        console.log("inside updates listener")
        const chart = Chart.getChart("myChart1");
        const parsedData = JSON.parse(event.data)
        chart.data.datasets[0].data = parsedData.data.results;
        chart.update();
    };}`,
		Call:       templ.SafeScript(`__templ_openEventConnection_6e34`, pollId),
		CallInline: templ.SafeScriptInline(`__templ_openEventConnection_6e34`, pollId),
	}
}

func copyUrl(url string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_copyUrl_ab83`,
		Function: `function __templ_copyUrl_ab83(url){const initialText = document.getElementById('copyURLButton').innerHTML;
    window.navigator.clipboard.writeText(url);

    document.getElementById('copyURLButton').innerHTML = "Copied!";
    setTimeout(function() {
        // Set button text back to what it was before it was changed to Copied!
        document.getElementById('copyURLButton').innerHTML = initialText;
    }, 500);}`,
		Call:       templ.SafeScript(`__templ_copyUrl_ab83`, url),
		CallInline: templ.SafeScriptInline(`__templ_copyUrl_ab83`, url),
	}
}

func loadChart(answers []string, pollResults []int, pollId string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_loadChart_61f0`,
		Function: `function __templ_loadChart_61f0(answers, pollResults, pollId){var fontSize = 12;
    if(window.innerWidth > 600) {
        fontSize = 16;
    };

    var data = {
        labels: answers,
        datasets: [{
            label: '# of Votes',
            data: pollResults,
            borderWidth: 1
        }]
    };

    var options = {
        maintainAspectRatio: true,
            scales: {
                y: {
                    beginAtZero: true,
                    grace: 5,
                    ticks: {
                        stepSize: 5
                }
            }
        }
    }

    Chart.defaults.font.size = fontSize;

    var myChart = new Chart("myChart1", {
        type: 'bar',
        data: data,
        options: options
    });}`,
		Call:       templ.SafeScript(`__templ_loadChart_61f0`, answers, pollResults, pollId),
		CallInline: templ.SafeScriptInline(`__templ_loadChart_61f0`, answers, pollResults, pollId),
	}
}

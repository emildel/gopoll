// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CreatorChartView(title string, answers []string, pollResults []int, pollId string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><div class=\"bg-[#FCFDFC] p-4\"><div class=\"flex justify-center items-center relative m-auto h-[50vh] max-w-[1200px] px-2 min-[601px]:px-4 shadow-lg shadow-slate-200 rounded\"><canvas id=\"myChart1\"></canvas></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = openEventConnection(pollId).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func PollCreator(title string, answers []string, pollResults []int, pollId string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var5 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = CreatorChartView(title, answers, pollResults, pollId).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Head("GoPoll | Your Poll").Render(templ.WithChildren(ctx, templ_7745c5c3_Var5), templ_7745c5c3_Buffer)
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
		Name: `__templ_openEventConnection_da89`,
		Function: `function __templ_openEventConnection_da89(pollId){const eventSource = new EventSource("/updateChart/" + pollId + "?stream="+pollId);

    eventSource.onopen = function() {
        console.log("event openned")
    };
    
    // eventSource.addEventListener(pollId, (event) => {
    //     console.log("inside updates listener")
    //     const parsedData = JSON.parse(event.data)
    //     let chart = Chart.getChart("myChart1")
    //     chart.data.datasets[0].data = parsedData.data.results;
    //     chart.update();
    // });

    eventSource.onmessage = (event) => {
        console.log("inside updates listener")
        const chart = Chart.getChart("myChart1");
        const parsedData = JSON.parse(event.data)
        chart.data.datasets[0].data = parsedData.data.results;
        chart.update();
    };}`,
		Call:       templ.SafeScript(`__templ_openEventConnection_da89`, pollId),
		CallInline: templ.SafeScriptInline(`__templ_openEventConnection_da89`, pollId),
	}
}

func loadChart(answers []string, pollResults []int, pollId string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_loadChart_3a3d`,
		Function: `function __templ_loadChart_3a3d(answers, pollResults, pollId){// async function subscribe(pollId) {
    //     const endpoint = window.location.origin + "/updateChart/" + pollId;

    //     const response = await fetch(endpoint, {
    //         headers: {
    //             'Accept': 'application/json'
    //         }
    //     });

    //     if (response.status != 200) {
    //         console.log(response.status);

    //         await new Promise(resolve => setTimeout(resolve, 1000));
    //         await subscribe(pollId);

    //     } else {
    //         const results = await response.json();

    //         const scores = results.pollResults.results.map(function(index) {
    //             return index;
    //         });

    //         myChart.data.datasets[0].data = scores;
    //         myChart.update();

    //         await subscribe(pollId);
    //     };
    // };

    // subscribe(pollId);

    var fontSize = 12;
    if(window.innerWidth > 600) {
        fontSize = 16;
    }

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
		Call:       templ.SafeScript(`__templ_loadChart_3a3d`, answers, pollResults, pollId),
		CallInline: templ.SafeScriptInline(`__templ_loadChart_3a3d`, answers, pollResults, pollId),
	}
}

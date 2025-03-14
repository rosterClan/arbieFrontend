import React , {useState, useEffect, useContext} from 'react';
import { useLoaderData, useParams } from 'react-router-dom';
import Chart, { elements } from 'chart.js/auto';
import zoomPlugin from 'chartjs-plugin-zoom';
import { Line } from "react-chartjs-2";

import './graph.css'

function getFormattedDateTime() {
    const now = new Date();
    const year = now.getUTCFullYear();
    const month = String(now.getUTCMonth() + 1).padStart(2, '0');
    const day = String(now.getUTCDate()).padStart(2, '0');
    const hours = String(now.getUTCHours()).padStart(2, '0');
    const minutes = String(now.getUTCMinutes()).padStart(2, '0');
    const seconds = String(now.getUTCSeconds()).padStart(2, '0');
    const milliseconds = String(now.getUTCMilliseconds()).padStart(3, '0');
    
    const microseconds = String(Math.floor(Math.random() * 1000)).padStart(3, '0');

    return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}.${milliseconds}${microseconds}Z`;
}

const EntrantChart = (props) => {
    Chart.register(zoomPlugin);
    const [entrant_data, set_entrant_data] = useState(props.entrant_data);
    const [price_data, set_price_data] = useState([]);
    const [graph_data, set_graph_data] = useState({labels: ['2024-08-10T00:00:00Z', '2024-08-10T01:00:00Z'],datasets: []});

    useEffect(() => {
        var local_price_data = entrant_data['Odds'];
        var test_data = {labels: [1,2,3,4,5,6,7], datasets: []};
        var x_axis = new Set();

        local_price_data.forEach((element,idx) => {

            element['Price_Fluctuations'].forEach((price_entry,idx) => {
                if (!(price_entry['Record_Time'] in x_axis)) {
                    let date_obj = new Date(price_entry['Record_Time']);
                    let pre_date = new Date(date_obj.getTime() - 5);
                    let post_date = new Date(date_obj.getTime() + 5);

                    x_axis.add(date_obj);
                    x_axis.add(pre_date);
                    x_axis.add(post_date);
                }
            })
        })
        x_axis = Array.from(x_axis);
        x_axis.push(new Date())
        x_axis.sort();

        var x_axis_idx = 0;
        var x_axis_idx_cache = {}
        x_axis.forEach((ele,idx) => {
            x_axis_idx_cache[ele] = x_axis_idx;
            x_axis_idx += 1;
        })

        local_price_data.forEach((platform,idx) => {
            var data = Array(x_axis.length).fill(null);
            console.log("hit", platform);
            platform['Price_Fluctuations'].forEach((price_instance,idx) => {
                var date_idx = new Date(price_instance['Record_Time']);
                data[x_axis_idx_cache[date_idx]] = price_instance['Odds'];
            })
            var smug_val = null;
            for (var idx = 0; idx < data.length; idx++) {
                if (smug_val == null && data[idx] != null) {
                    smug_val = data[idx];
                } else if (smug_val != null && data[idx] == null) {
                    data[idx] = smug_val;
                } else if (smug_val != null && data[idx] != null) {
                    smug_val = data[idx];
                } else {
                    data[idx] = 0;
                }
            }

            test_data['datasets'].push(
                {
                    label: platform['Platform_Name'],
                    data: data,
                    fill: false,
                    backgroundColor: "rgb(15, 107, 161)",
                    borderColor: "rgb(15, 107, 161)"
                }
            )

        })  

        x_axis.forEach((element,idx) => {
            element = element.toISOString();
        })
        test_data['labels'] = x_axis;
        console.log(test_data);
        set_graph_data(test_data);
    },[entrant_data])

    const options = {
        scales: {
            x: {
                display: false // Hides the x-axis labels
            }
        },
          plugins: {
            legend: {
                display: false
            },
            zoom: {
                pan: {
                    enabled: true,
                    mode: 'x'
                },
                zoom: {
                    wheel: {
                        enabled: true
                    },
                    pinch: {
                        enabled: true
                    },
                    mode: 'x'
                }
            }
        }
    };


    const data_2 = {
        labels: ['2024-08-10T00:00:00Z', '2024-08-10T01:00:00Z'],
        datasets: []
    };

    return (
      <div class='entrant_graph_wpr'>
        <div class='entrant_graph_header'>
            <p>{props.entrant_data["Entrant_Name"]}</p>
        </div>
        <div class='entrant_graph_body'>
            <Line data={graph_data} options={options}/>
        </div>
      </div>
    );
}

export default EntrantChart; 

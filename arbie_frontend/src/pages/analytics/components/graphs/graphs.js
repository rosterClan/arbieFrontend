import React , {useState, useEffect, useContext} from 'react';
import EntrantChart from '../graph/graph';

import './graphs.css'

const EntrantGraphs = (props) => {
    const [entrants, set_entrants] = useState([]);
    const [non_filtered, set_non_filtered] = useState([]);

    useEffect(() => {
        console.log(props.entrants);
        let non_screated_entrants = [];
        non_filtered.forEach((entrant,idx) => {
            if (entrant['Is_Scratched'] == 0) {
                non_screated_entrants.push(entrant);
            }
        })
        set_entrants(props.entrants);
    },[props.entrants]);
    //            
    return (
      <div class='entrant_line_graph_wpr'>


            {entrants != undefined ? (
                entrants.map((item, index) => (

                        <EntrantChart entrant_data={item}/>

                ))
            ) : (<></>)}


      </div>
    );
}

export default EntrantGraphs; 
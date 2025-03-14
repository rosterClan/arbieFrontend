import React , {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';

import './round_select.css'

import CounterTime from '../../../counter_time/counter_time';

const RoundSelect = (props) => {
  const [round, set_round] = useState(1);
  const [race_data,set_race_data] = useState(props.race);
  const navigate = useNavigate(); 

  useEffect(() => {
    set_round(props.round);
  },[props.round]);

  useEffect(() => {
    set_race_data(props.race_data);
  },[props.race_data]);

  const on_day_trigger = (element) => {
    console.log(element);
  }

  return (
    <div style={{height: props.height}} onClick={()=>navigate(`/race?track=${props.race_data["Track_Name"]}&round=${props.race_data["Round"]}&start_time=${props.race_data["Start_Time"]}`)} class={'round_select_container'}>
      <div class='padded_content'>
        
        <div class='vertical_divide'>
          <div class='title_wpr'>
            <p>R{round}</p>
          </div>
        </div>

        <div class='vertical_divide'>
          <CounterTime trigger_function={() => on_day_trigger(this)} target_time={new Date(props.race_data['Start_Time'])}/>
        </div>

      </div>
    </div>
  );
}

export default RoundSelect;

import React , {useState, useEffect} from 'react';

import './next_2_go_tile.css'
import {useNavigate} from 'react-router-dom';
import CounterTime from '../../counter_time/counter_time';

const Next2GoTile = (props) => {
  const navigate = useNavigate(); 
  const [name, set_name] = useState(props.race_data['Track_Name']);
  const [round, set_round] = useState(props.race_data['Round']);
  const [start_time, set_start_time] = useState(props.race_data['Start_Time']);
  const [display_setting, set_display_setting] = useState('');

  let title = name + " R" + round.toString()

  const remove_tile = () => {
    set_display_setting("none");
  }

  return (
    <div onClick={()=>navigate(`/race?track=${name}&round=${round}&start_time=${start_time}`)} style={{display:display_setting}} class={'next_2_tile_container ' + props.race_data['Start_time']}>
      <div class='next_2_tile_image'>
        <svg viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg"><path d="M509.8 332.5l-69.9-164.3c-14.9-41.2-50.4-71-93-79.2 18-10.6 46.3-35.9 34.2-82.3-1.3-5-7.1-7.9-12-6.1L166.9 76.3C35.9 123.4 0 238.9 0 398.8V480c0 17.7 14.3 32 32 32h236.2c23.8 0 39.3-25 28.6-46.3L256 384v-.7c-45.6-3.5-84.6-30.7-104.3-69.6-1.6-3.1-.9-6.9 1.6-9.3l12.1-12.1c3.9-3.9 10.6-2.7 12.9 2.4 14.8 33.7 48.2 57.4 87.4 57.4 17.2 0 33-5.1 46.8-13.2l46 63.9c6 8.4 15.7 13.3 26 13.3h50.3c8.5 0 16.6-3.4 22.6-9.4l45.3-39.8c8.9-9.1 11.7-22.6 7.1-34.4zM328 224c-13.3 0-24-10.7-24-24s10.7-24 24-24 24 10.7 24 24-10.7 24-24 24z"/></svg>
      </div>
      <div class='next_2_tile_details'>
        <div class='next_2_gp_ele title'>
          <p>{title}</p>
        </div>
        <div class='next_2_gp_ele time'>
          <CounterTime trigger_function={remove_tile} target_time={new Date(start_time)}/>
        </div>
      </div>
    </div>
  );
}

export default Next2GoTile;

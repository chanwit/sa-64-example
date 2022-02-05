import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import { UsersInterface } from "../models/IUser";
import { PlaylistsInterface } from "../models/IPlaylist";
import { ResolutionsInterface } from "../models/IResolution";
import { VideosInterface } from "../models/IVideo";
import { WatchVideoInterface } from "../models/IWatchVideo";

import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function WatchVideoCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [users, setUsers] = useState<UsersInterface[]>([]);
  const [videos, setVideos] = useState<VideosInterface[]>([]);
  const [resolutions, setResolutions] = useState<ResolutionsInterface[]>([]);
  const [playlists, setPlaylists] = useState<PlaylistsInterface>();
  const [watchVideo, setWatchVideo] = useState<Partial<WatchVideoInterface>>(
    {}
  );

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof watchVideo;
    setWatchVideo({
      ...watchVideo,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };

  const getUsers = async () => {
    fetch(`${apiUrl}/users`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setUsers(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getVideos = async () => {
    fetch(`${apiUrl}/videos`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setVideos(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getResolution = async () => {
    fetch(`${apiUrl}/resolutions`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setResolutions(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getPlaylist = async () => {
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/playlist/watched/user/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        watchVideo.PlaylistID = res.data.ID
        if (res.data) {
          setPlaylists(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getUsers();
    getVideos();
    getResolution();
    getPlaylist();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    let data = {
      ResolutionID: convertType(watchVideo.ResolutionID),
      PlaylistID: convertType(watchVideo.PlaylistID),
      VideoID: convertType(watchVideo.VideoID),
      WatchedTime: selectedDate,
    };

    console.log(data)

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/watch_videos`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true)
          setErrorMessage("")
        } else {
          console.log("บันทึกไม่ได้")
          setError(true)
          setErrorMessage(res.error)
        }
      });
  }

  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ: {errorMessage}
        </Alert>
      </Snackbar>
      <Paper className={classes.paper}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกการเข้าชมวีดีโอ
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วีดีโอ</p>
              <Select
                native
                value={watchVideo.VideoID}
                onChange={handleChange}
                inputProps={{
                  name: "VideoID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกวีดีโอ
                </option>
                {videos.map((item: VideosInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ความละอียด</p>
              <Select
                native
                value={watchVideo.ResolutionID}
                onChange={handleChange}
                inputProps={{
                  name: "ResolutionID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกความละอียด
                </option>
                {resolutions.map((item: ResolutionsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Value}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เพลย์ลิสต์</p>
              <Select
                native
                value={watchVideo.PlaylistID}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "PlaylistID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกเพลย์ลิสต์
                </option>
                <option value={playlists?.ID} key={playlists?.ID}>
                  {playlists?.Title}
                </option>

                {/* {playlists.map((item: PlaylistsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Title}
                  </option>
                ))} */}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="WatchedTime"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label="กรุณาเลือกวันที่และเวลา"
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/watch_videos"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default WatchVideoCreate;

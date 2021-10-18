import { PlaylistsInterface } from "./IPlaylist";
import { ResolutionsInterface } from "./IResolution";
import { VideosInterface } from "./IVideo";

export interface WatchVideoInterface {
  ID: number,
  WatchedTime: Date,
  ResolutionID: number,
  Resolution: ResolutionsInterface,
  PlaylistID: number,
  Playlist: PlaylistsInterface,
  VideoID: number,
  Video: VideosInterface,
}

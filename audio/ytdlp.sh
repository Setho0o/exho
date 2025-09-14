#!bin/sh
[ "$1" = "-h" -o "$1" = "--help" -o "$1" = "" ] && echo "
  format + dir + url
 
  url: youtube url, if you do a playlist url it will download everything.
  format: optional, defaults to wav. The audio formats exho supports are wav, flac, and mp3.
  dir: optional, this creates a dir to hold the songs in the music folder.  

  ex: ./ytdlp.sh wav https://youtube.com/... 
  
  As you should expect yt-dlp is required for this so install that.
  This is just a simple wrapper for the yt-dlp repo, go there for more options.

  https://github.com/yt-dlp/yt-dlp
" && exit

yt-dlp \
  -P "music/"$3\
  --parse-metadata "description:(?s)(?P<meta_comment>.+)" \
  --parse-metadata "%(webpage_url)s:%(url)s" \
  --parse-metadata "%(like_count)s:%(meta_likes)s" \
  --parse-metadata "%(view_count)s:%(meta_view)s" \
  --parse-metadata "%(average_rating)s:%(rating)s" \
  --parse-metadata "video::(?P<formats>)" \
  --replace-in-metadata "title,uploader" "[ -]" "_" \
  -o "%(title)s"\
  --embed-thumbnail \
  --embed-metadata --merge-output-format mkv --write-info-json \
  -x --audio-format $1 $4 $2 


audio_file_name="./sample"
default_audio_format="wav"

echo $audio_file_name | sed -E -n 's#^([0-9]+)-max-([0-9.]+)-mid-([0-9.]+)\.'"${default_audio_format}"'$#\2#p'

echo "... maximum amplitude: ${max_amplitude}"
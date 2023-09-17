#!/bin/bash

main_directory="$1 at $(date '+%a %b %e %T %Z %Y')"
mkdir -p "$main_directory"

facebook_profile="$2"
mkdir -p "$main_directory/about_me/personal"
echo "https://www.facebook.com/$facebook_profile" > "$main_directory/about_me/personal/facebook.txt"

linkedin_profile="$3"
mkdir -p "$main_directory/about_me/professional"
echo "https://www.linkedin.com/in/$linkedin_profile" > "$main_directory/about_me/professional/linkedin.txt"

mkdir -p "$main_directory/my_friends"
curl -o "$main_directory/my_friends/list_of_my_friends.txt" https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt

mkdir -p "$main_directory/my_system_info"
about_this_laptop_info=$(uname -a)
echo "My Username : $USER" > "$main_directory/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$main_directory/my_system_info/about_this_laptop.txt"

ping -c 3 google.com > "$main_directory/my_system_info/internet_connection.txt"

tree "$main_directory"

echo "Program telah selesai dijalankan"
#!/usr/bin/env bash

TIMESTAMP=$(date +%Y%m%d-%H%M%S)

while [[ true ]] ; do
  # waiting for modifications to config file or timeout after 300s
  inotifywait -t 300 -e modify $PATHTOYOURKEYSTOREFOLDER/keystore

# Copy files

sudo cp -r $PATHTOYOURKEYSTOREFOLDER/keystore $PATHTOYOURKEYSTOREBACKUPFOLDER/keystore-backup

echo "Folder copied"

# Zip folders

sudo chown -R $YOURUSER:$YOURUSER $PATHTOYOURKEYSTOREBACKUPFOLDER/keystore-backup

zip -rP $YOURZIPPASSWORD $PATHTOYOURKEYSTOREBACKUPFOLDER/keystore-backup/accounts-backups-${TIMESTAMP}.zip $PATHTOYOURKEYSTOREBACKUPFOLDER/backup/keystore-backup

echo "Folders zipped"

# Transfer files

rclone --bwlimit 100M sync -q $PATHTOYOURKEYSTOREBACKUPFOLDER/keystore-backup/accounts-backups-${TIMESTAMP}.zip  $YOURRCLONEDESTINATION

echo "File transferred"

# Remove folders and files

rm $PATHTOYOURKEYSTOREBACKUPFOLDER/keystore-backup/accounts-backups-${TIMESTAMP}.zip

rm -rf $PATHTOYOURKEYSTOREBACKUPFOLDER/keystore-backup/keystore

echo "Folders and files removed"

echo "Backup done at: $(date +%Y%m%d-%H%M%S)"

done

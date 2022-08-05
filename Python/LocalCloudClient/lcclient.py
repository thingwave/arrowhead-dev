#!/usr/bin/env python3
import sys
import argparse
import requests

def main(argv):
    print('Eclipse Arrowhead Local cloud HTTP(S) client tool\nCopyright 2022 ThingWave AB')
    parser = argparse.ArgumentParser()
    parser.add_argument('--cloud', type=str, required=False, help='Local clound certificate file')
    parser.add_argument('--cert', type=str, required=False, help='System certificate file')
    parser.add_argument('--key', type=str, required=False, help='System certificate key file')
    parser.add_argument('--insecure',  action="store_true", required=False, help='To run in INSECURE mode')
    parser.add_argument('--uri', type=str, required=True, help='URI to call')
    args = parser.parse_args()

    if args.insecure == True:
        try:
            r = requests.get(url=args.uri, params=None)
            print(r)
        except:
            print("Could not connect to: '" + args.uri + "'")
    

if __name__ == '__main__':
    main(sys.argv[1:])


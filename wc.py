import sys
import json
from collections import defaultdict

def count_visits(filename):

    with open(filename) as f:
        data = f.read()
    print "File read"
    return [1]

    json_data = json.loads(data)
    print "Json Parsed"

    items = json_data['register']

    ids = defaultdict(lambda : 0)

    for i, item in enumerate(items):
        key = item['status']['objectId']
        ids[key] += 1
        if i % 10000 == 0:
            print i

    return ids


if __name__ == '__main__':

    counts = count_visits(sys.argv[1])
    print len(counts)
    # for id, count in counts.iteritems():
    #     print id, count

# { "register": [

# ]
#}
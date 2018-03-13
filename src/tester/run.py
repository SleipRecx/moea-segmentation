import os

from PIL import Image
import numpy as np

optimalFolder = "Optimal_Segmentation_Files"  # you may have to specify the complete path
studentFolder = "Student_Segmentation_Files"  # you may have to specify the complete path
colorValueSlackRange = 40
blackValueThreshold = 100  # colors below 100 is black
pixelRangeCheck = 4
checkEightSurroundingPixels = True

# you need to download PIL:
#for python 3: python3 -m pip install Pillow
#for python 2: sudo pip install Pillow
#file reader: reads file name and returns 2d array

def readImage(filename):
    im = Image.open(filename)
    pix = im.load()
    width, height = im.size
    data = fixData(list(im.getdata()))
    #get pixel (x,y) by pixel_values[width*y+x]
    try:
        return np.array(data).reshape((width,height))
    except:
        raise Exception("Mo")

def fixData(data):
    #if data contains tuples, then replace every tuple with the first number in the tuple
    check = type(data[0])
    if check == tuple:
        for i in range(len(data)):
            data[i] = data[i][0]
    return data

def readTextFile(filename):
    width = 0
    height = 0
    im = np.array([])
    file = open(filename, "r")
    for line in file:
        imLine = line.strip().split(",")
        width = len(imLine)
        #print(width)
        for i in range(len(imLine)):
            imLine[i] = int(imLine[i])
        im = np.append(im, np.array(imLine))
        height += 1
    #print(height)
    try:
        im.reshape((width,height))
    except:
        raise Exception('\n\nSome error with the shape of the .txt image file \n\n')
    return im.reshape((width,height))




def readFilesFromFolder(directory, filter=None):
    allFiles = []
    for filename in os.listdir(directory):
        if not filter or filter in filename.upper():
            if filename.endswith(".jpg") or filename.endswith(".png"):
                filename = os.path.join(directory, filename)
                allFiles.append(readImage(filename))
            elif filename.endswith(".txt"):
                filename = os.path.join(directory, filename)
                allFiles.append(readTextFile(filename))
    return allFiles


def comparePics(studentPic, optimalSegmentPic):
	# for each pixel in studentPic, compare to corresponding pixel in optimalSegmentPic
	global colorValueSlackRange
	global checkEightSurroundingPixels
	global pixelRangeCheck
	width, height= studentPic.shape
	counter = 0 #counts the number of similar pics
	numberOfBlackPixels = 0
	for w in range(width):
		for h in range(height):
			#if any pixel nearby or at the same position is within the range, it counts as correct
			color1 = studentPic[w][h]
			color2 = optimalSegmentPic[w][h]
			if color1 < blackValueThreshold:
				#black color
				numberOfBlackPixels +=1
				if(color1 - colorValueSlackRange< color2  and color2 < colorValueSlackRange + color1):
					counter +=1
					continue
				elif checkEightSurroundingPixels:
					#check surroundings
					correctFound = False
					for w2 in range(w-pixelRangeCheck, w + pixelRangeCheck+1):
						if(correctFound):
							break
						for h2 in range(h - pixelRangeCheck, h + pixelRangeCheck+1):
							if(w2 >=0 and h2 >= 0 and w2 < width and h2 < height):

								color2 = optimalSegmentPic[w2][h2]
								if( color1 - colorValueSlackRange< color2  and color2 < colorValueSlackRange + color1):
									correctFound = True
									counter +=1
									break

	return float(counter)/float(max(numberOfBlackPixels,1))


if __name__ == '__main__':
    import sys

    if len(sys.argv) < 3:
        raise ValueError("TOO FEW ARGUMENTS!")

    optimalFiles = readFilesFromFolder(sys.argv[1], filter="GT")
    studentFiles = readFilesFromFolder(sys.argv[2])
    totalScore = 0
    for student in studentFiles:
        highestScore = 0
        for opt in optimalFiles:
            result1 = comparePics(opt, student)
            result2 = comparePics(student, opt)
            result = min(result1, result2)
            highestScore = max(highestScore, result)
        totalScore += highestScore
        a = highestScore * 100
        print("Score: %.2f" % (a) + "%")
    a = totalScore / len(studentFiles) * 100
    print("Total Average Score: %.2f" % a + "%")



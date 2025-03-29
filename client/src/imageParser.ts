// Image file class operation
class ImageFile {
  imgFile: File;

  constructor(imgFile: File) {
    this.imgFile = imgFile;
  }

  // Print out the byte contents of the file
  async displayFileContents() {
    const buffer = await this.imgFile.arrayBuffer();
    for (const b of new Uint8Array(buffer)) {
      console.log(b);
    }
  }
}

export class PNGImge extends ImageFile {
  constructor(pngFile: File) {
    super(pngFile);
  }

  private async fileToBuffer(): Promise<Uint8Array> {
    const buff = await this.imgFile.arrayBuffer();
    return new Uint8Array(buff);
  }

  // Parse the contents of the png file
  async parsePng() {
    try {
      // TODO: Parse the png file, if the structure is not a valid png format throw an error and catch it immedietly
      const bufferArray = await this.fileToBuffer();
      throw new Error('InvalidPNGByteStructureError');
    } catch (e) {
      if (e instanceof Error && e.message === 'InvalidPNGByteStructureError') {
        console.error('Byte structure of image is not png formated...\n');
        console.error(e);
      } else {
        console.error(e);
      }
    }
  }
}

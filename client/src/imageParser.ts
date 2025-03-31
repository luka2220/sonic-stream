// Image file class operation
class ImageFile {
  data: File;

  constructor(imgFile: File) {
    this.data = imgFile;

    if (!this.validateImageSize()) {
      console.error('image size too big, must be less than 500kb');
      throw new Error('InvalidImageSize');
    }
  }

  private validateImageSize(): boolean {
    return this.data.size < 500_000 ? true : false;
  }

  async displayFileContents() {
    const buffer = await this.data.arrayBuffer();
    for (const b of new Uint8Array(buffer)) {
      console.log(b);
    }
  }
}

export class PNGImge extends ImageFile {
  constructor(pngFile: File) {
    super(pngFile);
  }

  async validatePng() {
    try {
      const buffer = await this.data.arrayBuffer();
      const bufferArray = new Uint8Array(buffer);

      const expectedPNGSignature = [
        0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a,
      ];
      const actual = bufferArray.subarray(0, 8);

      for (let i = 0; i < expectedPNGSignature.length; i++) {
        if (actual[i] !== expectedPNGSignature[i]) {
          throw new Error('InvalidPNGByteStructureError');
        }
      }
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

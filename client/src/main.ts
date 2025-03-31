import { PNGImge } from './imageParser';

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <section class="mt-5">
    <h1 class="text-4xl">Sonic Stream Client</h1>
  </section>
  <section id="image-upload-section" class= "mt-10"> 
    <label for="img-file">Upload an image file <span>(PNG, JPEG, GIF, BMP, WebP)</span></label>
    <br><br>
    <input  class="block w-full max-w-sm text-sm text-white
         file:mr-4 file:py-2 file:px-4
         file:rounded file:border-0
         file:bg-blue-600 file:text-white
         file:hover:bg-blue-700
         file:cursor-pointer" type="file" id="img-file" name="img-file" accept="image/png, image/jpeg, image/gif, image/bmp, image/webp"/>
  </section>
`;

// listen for image file upload
// Currently only image size of 500kb
document
  .querySelector<HTMLInputElement>('#img-file')
  ?.addEventListener('input', async (e: Event) => {
    if (!(e.target instanceof HTMLInputElement)) {
      return;
    }

    const files = e.target.files;
    const imageFile = files?.item(0);
    if (!(imageFile instanceof File)) return;

    try {
      const pngFile = new PNGImge(imageFile);
      await pngFile.validatePng();
      console.log(pngFile.data.name);
      await sendImageToSonic(pngFile.data, 'jpeg');
    } catch (err) {
      if (err instanceof Error && err.message === 'InvalidImageSize') {
        // TODO: display an eror modal or message to the user
        // clear the input field
        e.target.value = '';
        addImageSizeErrorMessage('Max image size is 500kb');
        return;
      }
    }
  });

function addImageSizeErrorMessage(msg: string) {
  document.querySelector('#image-upload-section')?.insertAdjacentHTML(
    'beforeend',
    `
        <div id="imageUploadErrorMessage" class="mt-4"><p class="text-red-700 font-medium">${msg}</p></div>
        `,
  );

  // clean up image error message from the dom
  setTimeout(() => {
    document.querySelector('#imageUploadErrorMessage')?.remove();
  }, 2000);
}

async function sendImageToSonic(img: File, convert: string) {
  const data = new FormData();
  data.append('file', img);
  data.append('convert', convert);

  const response = await fetch('http://localhost:8080/api/image', {
    method: 'POST',
    body: data,
  });

  console.log(`sonic status = ${response.status}`);
  if (response.status === 200) {
    const json = await response.json();
    console.log(json);
  }
}

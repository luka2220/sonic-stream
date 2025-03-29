import { PNGImge } from './imageParser';

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <section class="mt-5">
    <h1 class="text-4xl">Sonic Stream Client</h1>
  </section>
  <section class= "mt-10"> 
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
document
  .querySelector<HTMLInputElement>('#img-file')
  ?.addEventListener('input', async (e: Event) => {
    if (!(e.target instanceof HTMLInputElement)) {
      return;
    }

    const files = e.target.files;
    const imageFile = files?.item(0);
    if (!(imageFile instanceof File)) return;
    const fOObj = new PNGImge(imageFile);
    await fOObj.displayFileContents();
  });

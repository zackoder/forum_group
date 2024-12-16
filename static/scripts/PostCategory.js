async function PostCategory() {
    let category=document.getElementById("category")
    try {
    const res =await fetch("http://localhost:8001/api/category/list" )
    const data=await res.json()
    data.forEach(catg => {
        
        category.innerHTML+=`
        <input type="checkbox" name="options" id="" value="${catg.Id}" data-name="${catg.Name}">${catg.Name}<br>
        `
        console.log(catg.Id);
    });
    } catch{
        console.log("erroure");
    }



}
PostCategory()




function SubmitPost() {
    // e.preventDefault();
    let categoryName= []
    let Title=document.getElementsByName("Title")
    let Content=document.getElementsByName("Content")
    category.addEventListener('click', (event) => {
        if (event.target.name === 'options' && event.target.type === 'checkbox') {
            categoryName.push(event.target.getAttribute('data-name')); 

        }
    });

    alert(Title.value);
    
    console.log(Content.value)
    

}
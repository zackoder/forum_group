async function PostCategory() {
    let category = document.getElementById("category")
    try {
        const res = await fetch("http://localhost:8001/api/category/list")
        const data = await res.json()
        data.forEach(catg => {

            category.innerHTML += `
        <input type="checkbox" name="options" id="" value="${catg.Name}" data-name="${catg.Name}">${catg.Name}<br>
        `
            console.log(catg.Id);
        });
    } catch {
        console.log("erroure");
    }



}
PostCategory()





//  function SubmitPost(e) {
//     e.preventDefault();
//     let isValidTitle = false
//     let isValidContent = false
//     let isvalidcheckboxes = false
//     let Title = document.getElementById("Title").value
//     let Content = document.getElementById("Content").value
//     let categoryName = []

//     let checkboxes = document.querySelectorAll('input[name="options"]:checked');

//     checkboxes.forEach(checkbox => {
//         categoryName.push(checkbox.getAttribute('data-name'));
//     });



//     if (Title === "") {
//         document.getElementById("errorTitle").innerHTML = "Title is null";
//         document.getElementById("errorTitle").style.color = "red";
//         isValidTitle = false;
//     } else {
//         document.getElementById("errorTitle").innerHTML = "";
//         isValidTitle = true;
//     }

//     if (Content === "") {
//         document.getElementById("errorContent").innerHTML = "Content is null";
//         document.getElementById("errorContent").style.color = "red";
//         isValidContent = false;
//     } else {
//         document.getElementById("errorContent").innerHTML = "";
//         isValidContent = true;
//     }


//     if (categoryName.length === 0) {
//         document.getElementById("errorcategory").innerHTML = "category is null";
//         document.getElementById("errorcategory").style.color = "red";
//         isvalidcheckboxes = false;
//     } else {
//         document.getElementById("errorcategory").innerHTML = "";
//         isvalidcheckboxes = true;
//     }



//     if (isvalidcheckboxes && isValidTitle && isValidContent) {
//         try {
//             const data =  res.json()
//             console.log(data);

//         } catch (error) {
//             console.log(error);

//         }
//     }

// }

document.getElementById("postForm").addEventListener("submit", async function (event) {
    event.preventDefault();  // Prevent the form from submitting

    // Declare validation flags
    let isValidTitle = true;
    let isValidContent = false;
    let isValidCheckboxes = false;

    // Get form values
    let Title = document.getElementById("Title").value;
    let Content = document.getElementById("Content").value;
    let categoryName = [];

    // Get selected checkboxes
    let checkboxes = document.querySelectorAll('input[name="options"]:checked');
    checkboxes.forEach(checkbox => {
        categoryName.push(checkbox.getAttribute('data-name'));
    });

    // Title validation
    // if (Title === "") {
    //     document.getElementById("errorTitle").innerHTML = "Title is required";
    //     document.getElementById("errorTitle").style.color = "red";
    //     isValidTitle = false;
    // } else {
    //     document.getElementById("errorTitle").innerHTML = "";
    //     isValidTitle = true;
    // }

    // Content validation
    if (Content === "") {
        document.getElementById("errorContent").innerHTML = "Content is required";
        document.getElementById("errorContent").style.color = "red";
        isValidContent = false;
    } else {
        document.getElementById("errorContent").innerHTML = "";
        isValidContent = true;
    }

    // Checkbox validation
    if (categoryName.length === 0) {
        document.getElementById("errorcategory").innerHTML = "Please select at least one category";
        document.getElementById("errorcategory").style.color = "red";
        isValidCheckboxes = false;
    } else {
        document.getElementById("errorcategory").innerHTML = "";
        isValidCheckboxes = true;
    }

    // If all validations are successful, submit the form
    if (isValidTitle && isValidContent && isValidCheckboxes) {
        try {
            const res = await fetch("http://localhost:8001/add-post", {
                method: "POST",
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                body: new URLSearchParams({
                    Title: Title,
                    Content: Content,
                    categoryName: categoryName
                }).toString()
            });

            if (res.ok) {
                alert('Post successfully submitted');

                // After the alert closes, redirect after a short delay (e.g., 500 milliseconds)
                setTimeout(function () {
                    window.location.href = res.url; // Redirect to the URL from the response
                }, 500);
            } else {
                alert("Failed to submit post");
                console.log(res);
            }
        } catch (error) {
            alert("Error: " + error.message);
        }
    }
});


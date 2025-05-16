using System.ComponentModel.DataAnnotations;

namespace WebCRUD.DTOs
{
    public class NewUserRequest
    {
        [Required(ErrorMessage = "Name is required")]
        [StringLength(100, ErrorMessage = "Name must be under 100 characters")]
        public string Name { get; set; } = string.Empty;

        [Required(ErrorMessage = "Email is required")]
        [EmailAddress(ErrorMessage = "Invalid email format")]
        public string Email { get; set; } = string.Empty;

        [Url(ErrorMessage = "Website must be a valid URL")]
        public string Website { get; set; } = string.Empty;
    }

    public class EditUserRequest
    {
        [Required]
        public int Id { get; set; }

        [Required(ErrorMessage = "Name is required")]
        [StringLength(100, ErrorMessage = "Name must be under 100 characters")]
        public string Name { get; set; } = string.Empty;

        [Required(ErrorMessage = "Email is required")]
        [EmailAddress(ErrorMessage = "Invalid email format")]
        public string Email { get; set; } = string.Empty;

        [Url(ErrorMessage = "Website must be a valid URL")]
        public string Website { get; set; } = string.Empty;
    }
}

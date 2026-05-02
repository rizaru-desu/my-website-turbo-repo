import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useForm } from "@tanstack/react-form";
import {
  Input,
  Label,
  Select,
  Textarea,
  Button,
} from "@repo/ui";
import {
  PublicNav,
  PublicFooter,
  SectionHeader,
  Sword,
  IconMail,
} from "../components";

export default function TestimonialFormPage() {
  const navigate = useNavigate();
  const [submitted, setSubmitted] = useState(false);

  const form = useForm({
    defaultValues: {
      name: "",
      role: "",
      company: "",
      message: "",
      rating: 5,
      relation: "CLIENT",
    },
    onSubmit: async ({ value }) => {
      // Mock submission
      await new Promise((resolve) => setTimeout(resolve, 1500));
      console.log("Submitted Testimonial:", value);
      setSubmitted(true);
    },
  });

  const relationOptions = [
    { value: "CLIENT", label: "CLIENT" },
    { value: "COLLEAGUE", label: "COLLEAGUE" },
    { value: "MENTOR", label: "MENTOR" },
    { value: "OTHER", label: "OTHER" },
  ];

  if (submitted) {
    return (
      <div className="public-shell">
        <PublicNav />
        <section className="section flex flex-col items-center justify-center min-h-[60vh] text-center">
          <div className="card accent-green p-10 max-w-[600px]">
            <Sword />
            <h2
              className="mt-6 mb-4 text-green"
              style={{ fontFamily: "var(--font-pixel)", fontSize: "1.5rem" }}
            >
              RECORD SAVED!
            </h2>
            <p
              className="text-muted mb-8 leading-relaxed"
              style={{ fontFamily: "var(--font-terminal)", fontSize: "1.1rem" }}
            >
              Thank you for your review. Your testimonial has been submitted for verification.
              It will appear in the REVIEWS.LOG once approved by the administrator.
            </p>
            <Button
              onClick={() => navigate("/")}
              className="hover-wiggle"
              style={{
                background: "var(--color-green)",
                color: "var(--color-bg)",
                boxShadow: "inset 0 3px 0 rgba(255, 255, 255, 0.35), inset 0 -3px 0 rgba(0, 0, 0, 0.35), 0 4px 0 #21222c"
              }}
            >
              RETURN TO BASE
            </Button>
          </div>
        </section>
        <PublicFooter />
      </div>
    );
  }

  return (
    <div className="public-shell">
      <PublicNav />

      <section className="section max-w-[800px] mx-auto">
        <SectionHeader
          kicker="~ NEW RECORD"
          title="ADD_TESTIMONIAL.EXE"
          subtitle="Leave a review for my character sheet."
        />

        <div className="card accent-purple">
          <form
            onSubmit={(e) => {
              e.preventDefault();
              e.stopPropagation();
              form.handleSubmit();
            }}
            className="flex flex-col gap-6"
          >
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              {/* Name */}
              <form.Field
                name="name"
                validators={{
                  onChange: ({ value }) =>
                    !value ? "Name is required" : undefined,
                }}
                children={(field) => (
                  <div className="flex flex-col gap-2">
                    <Label htmlFor={field.name}>&gt; NAME_ID</Label>
                    <Input
                      id={field.name}
                      name={field.name}
                      value={field.state.value}
                      onBlur={field.handleBlur}
                      onChange={(e) => field.handleChange(e.target.value)}
                      placeholder="Enter your name"
                    />
                    {field.state.meta.isTouched && field.state.meta.errors.length ? (
                      <em
                        className="text-red text-[0.7rem]"
                        style={{ fontFamily: "var(--font-terminal)" }}
                      >
                        {field.state.meta.errors.join(", ")}
                      </em>
                    ) : null}
                  </div>
                )}
              />

              {/* Role */}
              <form.Field
                name="role"
                validators={{
                  onChange: ({ value }) =>
                    !value ? "Role is required" : undefined,
                }}
                children={(field) => (
                  <div className="flex flex-col gap-2">
                    <Label htmlFor={field.name}>&gt; ROLE_CLASS</Label>
                    <Input
                      id={field.name}
                      name={field.name}
                      value={field.state.value}
                      onBlur={field.handleBlur}
                      onChange={(e) => field.handleChange(e.target.value)}
                      placeholder="e.g. CTO, Designer"
                    />
                    {field.state.meta.isTouched && field.state.meta.errors.length ? (
                      <em
                        className="text-red text-[0.7rem]"
                        style={{ fontFamily: "var(--font-terminal)" }}
                      >
                        {field.state.meta.errors.join(", ")}
                      </em>
                    ) : null}
                  </div>
                )}
              />

              {/* Company */}
              <form.Field
                name="company"
                children={(field) => (
                  <div className="flex flex-col gap-2">
                    <Label htmlFor={field.name}>&gt; GUILD_COMPANY</Label>
                    <Input
                      id={field.name}
                      name={field.name}
                      value={field.state.value}
                      onBlur={field.handleBlur}
                      onChange={(e) => field.handleChange(e.target.value)}
                      placeholder="e.g. Nebula Labs"
                    />
                  </div>
                )}
              />

              {/* Relation */}
              <form.Field
                name="relation"
                children={(field) => (
                  <div className="flex flex-col gap-2">
                    <Label htmlFor={field.name}>&gt; RELATION_TYPE</Label>
                    <Select
                      id={field.name}
                      name={field.name}
                      value={field.state.value}
                      onChange={(val) => field.handleChange(val)}
                      options={relationOptions}
                    />
                  </div>
                )}
              />
            </div>

            {/* Rating */}
            <form.Field
              name="rating"
              children={(field) => (
                <div className="flex flex-col gap-2">
                  <Label>&gt; RATING_LEVEL</Label>
                  <div className="flex gap-4">
                    {[1, 2, 3, 4, 5].map((star) => (
                      <button
                        key={star}
                        type="button"
                        onClick={() => field.handleChange(star)}
                        className={`text-2xl transition-transform hover:scale-110 ${
                          field.state.value >= star ? "text-yellow" : "text-line"
                        }`}
                        style={{ fontFamily: "var(--font-pixel)" }}
                      >
                        ★
                      </button>
                    ))}
                  </div>
                </div>
              )}
            />

            {/* Message */}
            <form.Field
              name="message"
              validators={{
                onChange: ({ value }) =>
                  !value
                    ? "Message is required"
                    : value.length < 10
                      ? "Message must be at least 10 characters"
                      : undefined,
              }}
              children={(field) => (
                <div className="flex flex-col gap-2">
                  <Label htmlFor={field.name}>&gt; LOG_MESSAGE</Label>
                  <Textarea
                    id={field.name}
                    name={field.name}
                    value={field.state.value}
                    onBlur={field.handleBlur}
                    onChange={(e) => field.handleChange(e.target.value)}
                    rows={5}
                    placeholder="Write your review here..."
                    className="resize-none"
                  />
                  {field.state.meta.isTouched && field.state.meta.errors.length ? (
                    <em
                      className="text-red text-[0.7rem]"
                      style={{ fontFamily: "var(--font-terminal)" }}
                    >
                      {field.state.meta.errors.join(", ")}
                    </em>
                  ) : null}
                </div>
              )}
            />

            <div className="flex justify-between items-center mt-4">
              <Button
                type="button"
                variant="ghost"
                onClick={() => navigate("/")}
                className="hover-wiggle"
              >
                CANCEL
              </Button>
              <form.Subscribe
                selector={(state) => [state.canSubmit, state.isSubmitting]}
                children={([canSubmit, isSubmitting]) => (
                  <Button
                    type="submit"
                    disabled={!canSubmit || isSubmitting}
                    className={`hover-wiggle ${
                      !canSubmit || isSubmitting ? "opacity-50 cursor-not-allowed" : ""
                    }`}
                  >
                    {isSubmitting ? "UPLOADING..." : <><IconMail /> SUBMIT_RECORD</>}
                  </Button>
                )}
              />
            </div>
          </form>
        </div>
      </section>

      <PublicFooter />
    </div>
  );
}
